package controllers

import (
	"net/http"

	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllToppingsHandler ...
// swagger:route GET /api/v1/toppings Toppings listToppings
//
// List all pizza toppings
//
// This will return the a list of toppings.
//
//
// Responses:
//   default: JSONResponse
//   200: JSONResponse
func GetAllToppingsHandler(w http.ResponseWriter, r *http.Request) {
	toppings, err := repository.AllToppings()
	helpers.RespondWithJSON(w, r, toppings, err)
}
