package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	repository "github.com/jayza/pizzaonthego/repositories"
)

/*
Pizza Route Handlers
*/

// GetOnePizzaHandler ...
func GetOnePizzaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pizzaID := mux.Vars(r)["id"]
	singlePizza := repository.OnePizza(pizzaID)
	json.NewEncoder(w).Encode(singlePizza)
}

// GetAllPizzasHandler ...
func GetAllPizzasHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pizzas := repository.AllPizzas()
	json.NewEncoder(w).Encode(pizzas)
}
