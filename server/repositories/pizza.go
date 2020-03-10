package pizza

import (
	"github.com/jayza/pizzaonthego/models"
	database "github.com/jayza/pizzaonthego/services"
)

// GetOne ...
func GetOne(id string) models.Pizza {
	db := database.Connect()
	const pizzaQuery = `
		SELECT p.id, p.name, p.price
		FROM pizzas as p
		WHERE p.id = ?
	`
	const pizzaOptionsQuery = `
		SELECT o.id, o.name, ot.name
		FROM options as o
		INNER JOIN pizzas_options as po ON ? = po.pizza_id
		INNER JOIN option_types as ot ON o.type_id = ot.id
		WHERE o.id = po.option_id
	`

	result, err := db.Query(pizzaQuery, id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	optionResult, err := db.Query(pizzaOptionsQuery, id)
	if err != nil {
		panic(err.Error())
	}
	defer optionResult.Close()

	var pizza models.Pizza

	for result.Next() {
		err := result.Scan(&pizza.ID, &pizza.Name, &pizza.Price)

		if err != nil {
			panic(err.Error())
		}

		for optionResult.Next() {
			var option models.Option

			err := optionResult.Scan(&option.ID, &option.Name, &option.Type)

			if err != nil {
				panic(err.Error())
			}

			pizza.Options = append(pizza.Options, option)
		}

	}

	return pizza
}

// GetAll ...
func GetAll() []models.Pizza {
	var db = database.Connect()
	const pizzaQuery = `
		SELECT p.id, p.name, p.price
		FROM pizzas as p
	`
	const pizzaOptionsQuery = `
		SELECT o.id, o.name, ot.name
		FROM options as o
		INNER JOIN pizzas_options as po ON ? = po.pizza_id
		INNER JOIN option_types as ot ON o.type_id = ot.id
		WHERE o.id = po.option_id
	`

	result, err := db.Query(pizzaQuery)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var pizzas []models.Pizza

	for result.Next() {
		var pizza models.Pizza

		err := result.Scan(&pizza.ID, &pizza.Name, &pizza.Price)

		if err != nil {
			panic(err.Error())
		}

		// Fetch options for the pizza.
		optionResult, err := db.Query(pizzaOptionsQuery, pizza.ID)

		if err != nil {
			panic(err.Error())
		}

		defer optionResult.Close()

		for optionResult.Next() {
			var option models.Option

			err := optionResult.Scan(&option.ID, &option.Name, &option.Type)

			if err != nil {
				panic(err.Error())
			}

			pizza.Options = append(pizza.Options, option)
		}

		pizzas = append(pizzas, pizza)
	}

	return pizzas
}
