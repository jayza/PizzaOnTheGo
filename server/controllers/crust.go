package controllers

import (
	"net/http"

	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllCrustsHandler ...
// swagger:route GET /api/v1/crusts Crusts listCrusts
//
// List all pizza crusts
//
// This will return the a list of crusts.
//
//
// Responses:
//   default: JSONResponse
//   200: JSONResponse
func GetAllCrustsHandler(w http.ResponseWriter, r *http.Request) {
	crusts, err := repository.AllCrusts()
	helpers.RespondWithJSON(w, r, crusts, err)
}
