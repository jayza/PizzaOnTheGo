package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// OnePizza ...
func OnePizza(id string) models.Pizza {
	const pizzaQuery = `
		SELECT p.id, p.name
		FROM product as p
		WHERE p.id = ?
	`
	// const pizzaOptionsQuery = `
	// 	SELECT o.id, o.name, ot.name
	// 	FROM pizza_option as o
	// 	INNER JOIN pizzas_pizza_options as po ON ? = po.pizza_id
	// 	INNER JOIN pizza_option_type as ot ON o.type_id = ot.id
	// 	WHERE o.id = po.pizza_option_id
	// `

	result, err := services.Db.Query(pizzaQuery, id)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var pizza models.Pizza

	for result.Next() {
		err := result.Scan(&pizza.ID, &pizza.Name)

		if err != nil {
			panic(err.Error())
		}

	}

	return pizza
}

// AllPizzas ...
func AllPizzas() []models.Pizza {
	const pizzaQuery = `
		SELECT p.id, p.name
		FROM product as p
		WHERE product_type_id = 1
	`
	// const pizzaOptionsQuery = `
	// 	SELECT o.id, o.name, ot.name
	// 	FROM pizza_option as o
	// 	INNER JOIN pizzas_pizza_options as po ON ? = po.pizza_id
	// 	INNER JOIN pizza_option_type as ot ON o.type_id = ot.id
	// 	WHERE o.id = po.pizza_option_id
	// `

	result, err := services.Db.Query(pizzaQuery)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var pizzas []models.Pizza

	for result.Next() {
		var pizza models.Pizza

		err := result.Scan(&pizza.ID, &pizza.Name)

		if err != nil {
			panic(err.Error())
		}

		// // Fetch options for the pizza.
		// optionResult, err := services.Db.Query(pizzaOptionsQuery, pizza.ID)

		// if err != nil {
		// 	panic(err.Error())
		// }

		// defer optionResult.Close()

		// for optionResult.Next() {
		// 	var option models.PizzaOption

		// 	err := optionResult.Scan(&option.ID, &option.Name, &option.Type)

		// 	if err != nil {
		// 		panic(err.Error())
		// 	}

		// 	switch option.Type {
		// 	case "Base":
		// 		pizza.Base = option
		// 	case "Topping":
		// 		pizza.Toppings = append(pizza.Toppings, option)
		// 	case "Crust":
		// 		pizza.Crust = option
		// 	case "Dough":
		// 		pizza.Dough = option
		// 	case "Size":
		// 		pizza.Size = option
		// 	}

		// pizza.Options = append(pizza.Options, option)
		// }

		pizzas = append(pizzas, pizza)
	}

	return pizzas
}
