package controllers

import (
	"net/http"

	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllBasesHandler ...
// swagger:route GET /api/v1/bases Bases listBases
//
// List all pizza bases
//
// This will return the a list of bases.
//
//
// Responses:
//   default: JSONResponse
//   200: JSONResponse
func GetAllBasesHandler(w http.ResponseWriter, r *http.Request) {
	bases, err := repository.AllBases()
	helpers.RespondWithJSON(w, r, bases, err)
}
