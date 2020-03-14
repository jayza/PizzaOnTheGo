package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// AllBases ...
func AllBases() (bases []*models.Ingredient, err error) {
	const baseQuery = `
		SELECT i.id, i.name, i.price, it.name FROM ingredient AS i
		INNER JOIN ingredient_type AS it ON it.id = i.ingredient_type_id
		WHERE ingredient_type_id = 1
	`

	result, err := services.Db.DB.Query(baseQuery)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var base *models.Ingredient = &models.Ingredient{}

		err := result.Scan(&base.ID, &base.Name, &base.Price, &base.Category)

		if err != nil {
			return nil, err
		}

		bases = append(bases, base)
	}

	return bases, nil
}

// OneBaseForPizza fetches the base of the pizza.
func OneBaseForPizza(pizzaID int) (base *models.Ingredient, err error) {
	const baseQuery = `
		SELECT i.id, i.name, i.price FROM ingredient AS i
		INNER JOIN product_ingredients AS pi ON pi.ingredient_id = i.id
		WHERE ingredient_type_id = 1
		AND pi.id = ?
	`

	base = &models.Ingredient{}

	err = services.Db.Find(baseQuery, services.Db.Params(pizzaID), services.Db.Fields(&base.ID, &base.Name, &base.Price))

	if err != nil {
		return nil, err
	}

	return base, nil
}
