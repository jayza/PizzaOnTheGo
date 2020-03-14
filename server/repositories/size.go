package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// AllSizes ...
func AllSizes() (sizes []*models.ProductSize, err error) {
	const sizeQuery = `
		SELECT ps.id, ps.name, ps.price FROM product_size AS ps
		WHERE product_type_id = 1
	`

	result, err := services.Db.DB.Query(sizeQuery)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var size *models.ProductSize = &models.ProductSize{}

		err := result.Scan(&size.ID, &size.Name, &size.Price)

		if err != nil {
			return nil, err
		}

		sizes = append(sizes, size)
	}

	return sizes, nil
}

// OneSize fetches a product size.
func OneSize(sizeID int) (size *models.ProductSize, err error) {
	const sizeQuery = `
		SELECT ps.id, ps.name, ps.price FROM product_size AS ps
		WHERE ps.product_type_id = 1
		AND ps.id = ?
	`

	size = &models.ProductSize{}
	err = services.Db.Find(sizeQuery, services.Db.Params(sizeID), services.Db.Fields(&size.ID, &size.Name, &size.Price))

	if err != nil {
		return nil, err
	}

	return size, nil
}
