package routers

import (
	"github.com/gorilla/mux"
	"github.com/jayza/pizzaonthego/controllers"
)

// GetRoutes ...
func GetRoutes() *mux.Router {
	// Create Server and Route Handlers
	routes := mux.NewRouter().StrictSlash(false)

	api := routes.PathPrefix("/api/v1").Subrouter()
	// Routes
	// routes.HandleFunc("", controllers.HomeHandler)
	api.HandleFunc("/pizzas", controllers.GetAllPizzasHandler).Methods("GET")
	api.HandleFunc("/pizzas/{id}", controllers.GetOnePizzaHandler).Methods("GET")

	return routes
}
