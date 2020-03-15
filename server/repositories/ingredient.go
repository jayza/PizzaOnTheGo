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
