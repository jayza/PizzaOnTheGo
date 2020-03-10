package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	pizza "github.com/jayza/pizzaonthego/repositories"
)

/*
 Miscellaneous Handlers
*/

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

/*
Pizza Route Handlers
*/

// GetOnePizzaHandler ...
func GetOnePizzaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pizzaID := mux.Vars(r)["id"]
	singlePizza := pizza.GetOne(pizzaID)
	json.NewEncoder(w).Encode(singlePizza)
}

// GetAllPizzasHandler ...
func GetAllPizzasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pizzas := pizza.GetAll()
	json.NewEncoder(w).Encode(pizzas)
}

/*
Topping Route Handler
*/

// GetAllToppingsHandler ...
// func GetAllToppingsHandler(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(toppings)
// }
