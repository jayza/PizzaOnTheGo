package controllers

import (
	"net/http"

	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllToppingsHandler ...
func GetAllToppingsHandler(w http.ResponseWriter, r *http.Request) {
	toppings, err := repository.AllToppings()
	helpers.RespondWithJSON(w, r, toppings, err)
}
