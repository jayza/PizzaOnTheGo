package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jayza/pizzaonthego/models"
)

type allPizzas []models.Pizza
type allToppings []models.Topping

var pizzas = allPizzas{
	{
		ID:    "1",
		Name:  "Margherita",
		Price: 9000,
	},
}

var toppings = allToppings{
	{
		ID:   "1",
		Name: "Tomato sauce",
	},
	{
		ID:   "2",
		Name: "Mozzarella",
	},
}

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
	pizzaID := mux.Vars(r)["id"]

	for _, singlePizza := range pizzas {
		if singlePizza.ID == pizzaID {
			json.NewEncoder(w).Encode(singlePizza)
		}
	}
}

// GetAllPizzasHandler ...
func GetAllPizzasHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pizzas)
}

/*
Topping Route Handler
*/

// GetAllToppingsHandler ...
func GetAllToppingsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(toppings)
}
