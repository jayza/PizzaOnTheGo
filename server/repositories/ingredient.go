package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// AllIngredientsForLineItem ...
func AllIngredientsForLineItem(lineItemID int) (ingredients []*models.Ingredient, err error) {
	const ingredientQuery = `
		SELECT i.id, i.name, i.price, it.name FROM ingredient AS i
		JOIN product_extra_ingredients AS pei ON pei.ingredient_id = i.id
		JOIN ingredient_type AS it ON it.id = i.ingredient_type_id
		WHERE pei.product_line_item_id = ?
	`

	result, err := services.Db.DB.Query(ingredientQuery, lineItemID)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var ingredient *models.Ingredient = &models.Ingredient{}

		err := result.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Price, &ingredient.Category)

		if err != nil {
			return nil, err
		}

		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

// AllIngredientsForPizza ...
func AllIngredientsForPizza(pizzaID int) (toppings []*models.Ingredient, base *models.Ingredient, dough *models.Ingredient, err error) {
	const ingredientQuery = `
		SELECT i.id, i.name, i.price, it.id, ic.name FROM ingredient AS i
		JOIN product_ingredients AS pi ON pi.ingredient_id = i.id
		JOIN ingredient_type AS it ON it.id = i.ingredient_type_id
		JOIN ingredient_category AS ic ON ic.id = i.ingredient_category_id
		WHERE pi.product_id = ?
	`

	result, err := services.Db.DB.Query(ingredientQuery, pizzaID)

	if err != nil {
		return nil, nil, nil, err
	}

	defer result.Close()

	for result.Next() {
		var ingredientType int
		var ingredient *models.Ingredient = &models.Ingredient{}

		err := result.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Price, &ingredientType, &ingredient.Category)

		if err != nil {
			return nil, nil, nil, err
		}

		switch ingredientType {
		case 1:
			base = ingredient
		case 2:
			toppings = append(toppings, ingredient)
		case 3:
			dough = ingredient
		}
	}

	return toppings, base, dough, nil
}
