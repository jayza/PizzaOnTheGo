package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jayza/pizzaonthego/helpers"
	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetOnePizzaHandler ...
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
func GetAllPizzasHandler(w http.ResponseWriter, r *http.Request) {
	pizzas, err := repository.AllPizzas()
	helpers.RespondWithJSON(w, r, pizzas, err)
}

// GetAllToppingsForPizzaHandler ...
func GetAllToppingsForPizzaHandler(w http.ResponseWriter, r *http.Request) {
	pizzaID, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		helpers.RespondWithError(w, r, 400, err)
		return
	}

	toppings, err := repository.AllToppingsForPizza(pizzaID)
	helpers.RespondWithJSON(w, r, toppings, err)
}
