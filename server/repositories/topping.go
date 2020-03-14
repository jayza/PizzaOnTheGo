package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// AllToppings ...
func AllToppings() (toppings []*models.Ingredient, err error) {
	const toppingQuery = `
		SELECT i.id, i.name, i.price, it.name FROM ingredient AS i
		INNER JOIN ingredient_type AS it ON it.id = i.ingredient_type_id
		WHERE i.ingredient_type_id = 2
	`

	result, err := services.Db.DB.Query(toppingQuery)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var topping *models.Ingredient = &models.Ingredient{}

		err := result.Scan(&topping.ID, &topping.Name, &topping.Price, &topping.Category)

		if err != nil {
			return nil, err
		}

		toppings = append(toppings, topping)
	}

	return toppings, nil
}

// AllToppingsByCategory gets all Toppings by Category ID
func AllToppingsByCategory(categoryID int) (toppings []*models.Ingredient, err error) {
	const toppingQuery = `
		SELECT i.id, i.name, i.price FROM ingredient AS i
		WHERE i.ingredient_type_id = 2
		AND i.ingredient_category_id = ?
	`

	result, err := services.Db.DB.Query(toppingQuery, categoryID)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var topping *models.Ingredient = &models.Ingredient{}

		err := result.Scan(&topping.ID, &topping.Name, &topping.Price, &topping.Category)

		if err != nil {
			return nil, err
		}

		toppings = append(toppings, topping)
	}

	return toppings, nil
}

// AllToppingsForPizza ...
func AllToppingsForPizza(pizzaID int) (toppings []*models.Ingredient, err error) {
	const toppingQuery = `
		SELECT i.id, i.name, i.price FROM ingredient AS i
		INNER JOIN product_ingredients AS pi ON pi.ingredient_id = i.id
		WHERE i.ingredient_type_id = 2
		AND pi.product_id = ?
	`

	result, err := services.Db.DB.Query(toppingQuery, pizzaID)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var topping *models.Ingredient = &models.Ingredient{}

		err := result.Scan(&topping.ID, &topping.Name, &topping.Price, &topping.Category)

		if err != nil {
			return nil, err
		}

		toppings = append(toppings, topping)
	}

	return toppings, nil
}
