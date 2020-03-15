package controllers

import (
	"net/http"

	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllDoughsHandler ...
// swagger:route GET /api/v1/doughs Doughs listDoughs
//
// List all pizza doughs
//
// This will return the a list of doughs.
//
//
// Responses:
//   default: JSONResponse
//   200: JSONResponse
func GetAllDoughsHandler(w http.ResponseWriter, r *http.Request) {
	doughs, err := repository.AllDoughs()
	helpers.RespondWithJSON(w, r, doughs, err)
}
