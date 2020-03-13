package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// AllBases ...
func AllBases() []models.Ingredient {
	const baseQuery = `
		SELECT i.id, i.name, i.price, it.name FROM ingredient AS i
		INNER JOIN ingredient_type AS it ON it.id = i.ingredient_type_id
		WHERE ingredient_type_id = 1
	`

	result, err := services.Db.Query(baseQuery)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var bases []models.Ingredient

	for result.Next() {
		var base models.Ingredient

		err := result.Scan(&base.ID, &base.Name, &base.Price, &base.Category)

		if err != nil {
			panic(err.Error())
		}

		bases = append(bases, base)
	}

	return bases
}
