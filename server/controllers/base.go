package controllers

import (
	"net/http"

	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllBasesHandler ...
func GetAllBasesHandler(w http.ResponseWriter, r *http.Request) {
	bases, err := repository.AllBases()
	helpers.RespondWithJSON(w, r, bases, err)
}
