package controllers

import (
	"net/http"

	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllCrustsHandler ...
func GetAllCrustsHandler(w http.ResponseWriter, r *http.Request) {
	crusts, err := repository.AllCrusts()
	helpers.RespondWithJSON(w, r, crusts, err)
}
