package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jayza/pizzaonthego/middlewares"
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
	"github.com/stretchr/testify/assert"
)

// Copied from https://github.com/gorilla/mux#testing-handlers
// But here we also check if the database can be pinged.
func TestHealthCheckHandler(t *testing.T) {
	db := services.NewDB(models.Env{Mock: true, T: t})
	defer db.DB.Close()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "{\"data\":{\"systemStatus\":true}}\n"
	assert.Equal(t, expected, rr.Body.String())
}

func TestAuthMiddlewareForbidden(t *testing.T) {
	routes := mux.NewRouter().StrictSlash(false)

	amw := middlewares.AuthMiddleware{}

	routes.Use(amw.Middleware)
	routes.HandleFunc("/healthz", HealthCheckHandler).Methods("GET")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	routes.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != 403 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, 403)
	}

	// Check the response body is what we expect.
	expected := "{\"errors\":{\"message\":\"Forbidden\",\"code\":403}}\n"
	assert.Equal(t, expected, rr.Body.String())
}

func TestAuthMiddlewareAuthenticated(t *testing.T) {
	db := services.NewDB(models.Env{Mock: true, T: t})
	defer db.DB.Close()

	routes := mux.NewRouter().StrictSlash(false)

	amw := middlewares.AuthMiddleware{}

	routes.Use(amw.Middleware)
	routes.HandleFunc("/healthz", HealthCheckHandler).Methods("GET")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/healthz?loggedInAs=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	routes.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "{\"data\":{\"systemStatus\":true}}\n"
	assert.Equal(t, expected, rr.Body.String())
}
