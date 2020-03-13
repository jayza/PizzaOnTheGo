package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// AllDoughs ...
func AllDoughs() []models.Ingredient {
	const doughQuery = `
		SELECT i.id, i.name, i.price, it.name FROM ingredient AS i
		INNER JOIN ingredient_type it ON i.ingredient_type_id = it.id
		WHERE ingredient_type_id = 3
	`

	result, err := services.Db.Query(doughQuery)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var doughs []models.Ingredient

	for result.Next() {
		var dough models.Ingredient

		err := result.Scan(&dough.ID, &dough.Name, &dough.Price, &dough.Category)

		if err != nil {
			panic(err.Error())
		}

		doughs = append(doughs, dough)
	}

	return doughs
}
