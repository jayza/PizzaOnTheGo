package controllers

import (
	"net/http"

	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllSizesHandler ...
func GetAllSizesHandler(w http.ResponseWriter, r *http.Request) {
	sizes, err := repository.AllSizes()
	helpers.RespondWithJSON(w, r, sizes, err)
}
