package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// AllCrusts ...
func AllCrusts() []models.ProductVariation {
	const crustQuery = `
		SELECT pv.id, pv.name, pv.price FROM product_variation AS pv
		WHERE product_type_id = 1
	`

	result, err := services.Db.Query(crustQuery)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var crusts []models.ProductVariation

	for result.Next() {
		var crust models.ProductVariation

		err := result.Scan(&crust.ID, &crust.Name, &crust.Price)

		if err != nil {
			panic(err.Error())
		}

		crusts = append(crusts, crust)
	}

	return crusts
}
