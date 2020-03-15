package controllers

import (
	"net/http"

	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllSizesHandler ...
// swagger:route GET /api/v1/sizes PizzaSizes listSizes
//
// List all pizza sizes
//
// This will return the a list of sizes.
//
//
// Responses:
//   default: JSONResponse
//   200: JSONResponse
func GetAllSizesHandler(w http.ResponseWriter, r *http.Request) {
	sizes, err := repository.AllSizes()
	helpers.RespondWithJSON(w, r, sizes, err)
}
