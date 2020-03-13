package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// AllSizes ...
func AllSizes() []models.ProductVariation {
	const sizeQuery = `
		SELECT ps.id, ps.name, ps.price FROM product_size AS ps
		WHERE product_type_id = 1
	`

	result, err := services.Db.Query(sizeQuery)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var sizes []models.ProductVariation

	for result.Next() {
		var size models.ProductVariation

		err := result.Scan(&size.ID, &size.Name, &size.Price)

		if err != nil {
			panic(err.Error())
		}

		sizes = append(sizes, size)
	}

	return sizes
}
