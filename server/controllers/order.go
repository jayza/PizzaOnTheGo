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
// swagger:route GET /api/v1/orders/:id Orders findOrder
//
// Finds order by order id parameter.
//
// This will return the entire order, complete with
// line items and shipping information.
//
//
// Responses:
//   default: JSONResponse
//   200: JSONResponse
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
// swagger:route POST /api/v1/orders Orders addOrder
//
// Creates a new order
//
// This will return the entire order, complete with
// line items and shipping information after creation.
//
//
//
//     Responses:
//       default: JSONResponse
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
