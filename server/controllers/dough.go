package controllers

import (
	"net/http"

	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllDoughsHandler ...
func GetAllDoughsHandler(w http.ResponseWriter, r *http.Request) {
	doughs, err := repository.AllDoughs()
	helpers.RespondWithJSON(w, r, doughs, err)
}
