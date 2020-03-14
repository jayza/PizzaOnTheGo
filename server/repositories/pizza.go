package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// OnePizza ...
func OnePizza(id string) (pizza *models.Pizza, err error) {
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

	result, err := services.Db.DB.Query(pizzaQuery, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	pizza = &models.Pizza{}

	for result.Next() {
		err := result.Scan(&pizza.ID, &pizza.Name)

		if err != nil {
			return nil, err
		}

		loadToppings, err := AllToppingsForPizza(pizza.ID)
		if loadToppings != nil {
			pizza.Toppings = loadToppings
		}
	}

	return pizza, nil
}

// AllPizzas ...
func AllPizzas() (pizzas []*models.Pizza, err error) {
	const pizzaQuery = `
		SELECT p.id, p.name
		FROM product as p
		WHERE product_type_id = 1
	`

	result, err := services.Db.DB.Query(pizzaQuery)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var pizza *models.Pizza = &models.Pizza{}

		err := result.Scan(&pizza.ID, &pizza.Name)

		if err != nil {
			return nil, err
		}

		pizzas = append(pizzas, pizza)
	}

	return pizzas, nil
}
