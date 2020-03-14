package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// AllDoughs ...
func AllDoughs() (doughs []*models.Ingredient, err error) {
	const doughQuery = `
		SELECT i.id, i.name, i.price, it.name FROM ingredient AS i
		INNER JOIN ingredient_type it ON i.ingredient_type_id = it.id
		WHERE ingredient_type_id = 3
	`

	result, err := services.Db.DB.Query(doughQuery)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var dough *models.Ingredient = &models.Ingredient{}

		err := result.Scan(&dough.ID, &dough.Name, &dough.Price, &dough.Category)

		if err != nil {
			return nil, err
		}

		doughs = append(doughs, dough)
	}

	return doughs, nil
}

// OneDoughForPizza fetches the dough of the pizza.
func OneDoughForPizza(pizzaID string) (dough *models.Ingredient, err error) {
	const doughQuery = `
		SELECT i.id, i.name, i.price FROM ingredient AS i
		INNER JOIN product_ingredients AS pi ON pi.ingredient_id = i.id
		WHERE ingredient_type_id = 3
		AND pi.id = ?
	`

	dough = &models.Ingredient{}

	err = services.Db.Find(doughQuery, services.Db.Params(pizzaID), services.Db.Fields(&dough.ID, &dough.Name, &dough.Price))
	if err != nil {
		return nil, err
	}

	return dough, nil
}
