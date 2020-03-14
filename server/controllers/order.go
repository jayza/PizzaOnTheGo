package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetOneOrderHandler ...
func GetOneOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := mux.Vars(r)["id"]
	order, err := repository.OneOrder(orderID)
	helpers.RespondWithJSON(w, r, order, err)
}
