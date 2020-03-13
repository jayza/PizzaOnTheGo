package controllers

import (
	"encoding/json"
	"net/http"

	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllToppingsHandler ...
func GetAllToppingsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	toppings := repository.AllToppings()
	json.NewEncoder(w).Encode(toppings)
}
