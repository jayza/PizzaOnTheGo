package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// AllToppings ...
func AllToppings() []models.Ingredient {
	const toppingQuery = `
		SELECT i.id, i.name, i.price, it.name FROM ingredient AS i
		INNER JOIN ingredient_type AS it ON it.id = i.ingredient_type_id
		WHERE ingredient_type_id = 2
	`

	result, err := services.Db.Query(toppingQuery)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var toppings []models.Ingredient

	for result.Next() {
		var topping models.Ingredient

		err := result.Scan(&topping.ID, &topping.Name, &topping.Price, &topping.Category)

		if err != nil {
			panic(err.Error())
		}

		toppings = append(toppings, topping)
	}

	return toppings
}
