package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jayza/pizzaonthego/helpers"
	"github.com/jayza/pizzaonthego/models"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetOneOrderHandler ...
func GetOneOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		helpers.RespondWithError(w, r, 400, err)
		return
	}

	order, err := repository.OneOrder(orderID)
	helpers.RespondWithJSON(w, r, order, err)
}

// CreateOrderHandler ...
func CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		helpers.RespondWithError(w, r, 400, err)
		return
	}

	o := models.Order{}
	json.Unmarshal(body, &o)

	order, err := repository.CreateOrder(o)

	helpers.RespondWithJSON(w, r, order, err)
}
