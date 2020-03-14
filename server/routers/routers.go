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

	// Pizzas
	api.HandleFunc("/pizzas", controllers.GetAllPizzasHandler).Methods("GET")
	api.HandleFunc("/pizzas/{id}", controllers.GetOnePizzaHandler).Methods("GET")
	pizza := api.PathPrefix("/pizzas/{id}").Subrouter()
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
	api.HandleFunc("/orders/{id}", controllers.GetOneOrderHandler).Methods("GET")

	return routes
}
