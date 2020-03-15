package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetOnePizzaHandler ...
// swagger:route GET /api/v1/pizzas/{id} Pizzas findPizza
//
// Find pizza by ID
//
// This will return the complete pizza.
//
//
// Responses:
//   default: JSONResponse
//   200: JSONResponse
func GetOnePizzaHandler(w http.ResponseWriter, r *http.Request) {
	pizzaID, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		helpers.RespondWithError(w, r, 400, err)
		return
	}

	pizza, err := repository.OnePizza(pizzaID)
	helpers.RespondWithJSON(w, r, pizza, err)
}

// GetAllPizzasHandler ...
// swagger:route GET /api/v1/pizzas Pizzas listPizzas
//
// List all non-custom pizzas
//
// This will return the entire order, complete with
// line items and shipping information.
//
//
// Responses:
//   default: JSONResponse
//   200: JSONResponse
func GetAllPizzasHandler(w http.ResponseWriter, r *http.Request) {
	pizzas, err := repository.AllPizzas()
	helpers.RespondWithJSON(w, r, pizzas, err)
}

// GetAllToppingsForPizzaHandler ...
// swagger:route GET /api/v1/pizzas/{id}/toppings Pizzas findPizzaToppings
//
// List toppings for Pizza
//
// This will return all the toppings on a Pizza specified by the
// given id parameter.
//
//
// Responses:
//   default: JSONResponse
//   200: JSONResponse
func GetAllToppingsForPizzaHandler(w http.ResponseWriter, r *http.Request) {
	pizzaID, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		helpers.RespondWithError(w, r, 400, err)
		return
	}

	toppings, err := repository.AllToppingsForPizza(pizzaID)
	helpers.RespondWithJSON(w, r, toppings, err)
}
