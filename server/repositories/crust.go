package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// AllCrusts ...
func AllCrusts() (crusts []*models.ProductVariation, err error) {
	const crustQuery = `
		SELECT pv.id, pv.name, pv.price FROM product_variation AS pv
		WHERE product_type_id = 1
	`

	result, err := services.Db.DB.Query(crustQuery)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var crust *models.ProductVariation = &models.ProductVariation{}

		err := result.Scan(&crust.ID, &crust.Name, &crust.Price)

		if err != nil {
			return nil, err
		}

		crusts = append(crusts, crust)
	}

	return crusts, nil
}

// OneCrust fetches a crust.
func OneCrust(crustID string) (crust *models.ProductVariation, err error) {
	const crustQuery = `
	SELECT pv.id, pv.name, pv.price FROM product_variation AS pv
	WHERE product_type_id = 1
	AND pv.id = ?
	`

	crust = &models.ProductVariation{}

	err = services.Db.Find(crustQuery, services.Db.Params(crustID), services.Db.Fields(&crust.ID, &crust.Name, &crust.Price))

	if err != nil {
		return nil, err
	}

	return crust, nil
}
