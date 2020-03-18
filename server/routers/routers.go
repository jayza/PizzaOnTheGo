package routers

import (
	"errors"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jayza/pizzaonthego/controllers"
	"github.com/jayza/pizzaonthego/errorshandler"
	"github.com/jayza/pizzaonthego/helpers"
	"github.com/jayza/pizzaonthego/middlewares"
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// GetRoutes ...
func GetRoutes() *mux.Router {
	// Create Server and Route Handlers
	routes := mux.NewRouter().StrictSlash(false)

	routes.HandleFunc("/healthz", HealthCheckHandler).Methods("GET")

	routes.PathPrefix("/public/receipts/").Handler(http.StripPrefix("/public/receipts/", http.FileServer(http.Dir(os.Getenv("RECEIPT_FILE_DIRECTORY")))))

	api := routes.PathPrefix("/api/v1").Subrouter()

	amw := middlewares.AuthMiddleware{}

	// Pizzas
	api.HandleFunc("/pizzas", controllers.GetAllPizzasHandler).Methods("GET")
	api.HandleFunc("/pizzas/{id:[0-9]+}", controllers.GetOnePizzaHandler).Methods("GET")

	pizza := api.PathPrefix("/pizzas/{id:[0-9]+}").Subrouter()
	pizza.HandleFunc("/toppings", controllers.GetAllToppingsForPizzaHandler).Methods("GET")

	//Toppings
	api.HandleFunc("/toppings", controllers.GetAllToppingsHandler).Methods("GET")

	//Doughs
	api.HandleFunc("/doughs", controllers.GetAllDoughsHandler).Methods("GET")

	//Bases
	api.HandleFunc("/bases", controllers.GetAllBasesHandler).Methods("GET")

	//Crusts
	api.HandleFunc("/crusts", controllers.GetAllCrustsHandler).Methods("GET")

	//Sizes
	api.HandleFunc("/sizes", controllers.GetAllSizesHandler).Methods("GET")

	//Orders
	order := api.PathPrefix("/orders").Subrouter()
	order.Use(amw.Middleware)

	order.HandleFunc("/{id:[0-9]+}", controllers.GetOneOrderHandler).Methods("GET")
	order.HandleFunc("", controllers.CreateOrderHandler).Methods("POST", "OPTIONS")
	order.HandleFunc("/{id:[0-9]+}/receipt", controllers.DownloadOrderReceipt).Methods("GET")

	aomw := middlewares.AllowOriginMiddleware{}
	api.Use(mux.CORSMethodMiddleware(api))
	api.Use(aomw.AllowOriginMiddleware)

	return routes
}

// HealthCheckHandler ...
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	var e error = nil
	var health models.Health = models.Health{Health: true}

	if err := services.Db.DB.Ping(); err != nil {
		e = errorshandler.HandleErrorCode(200, errors.New("could not reach database connection"))
		health.Health = false
	}

	helpers.RespondWithJSON(w, r, health, e)
}
