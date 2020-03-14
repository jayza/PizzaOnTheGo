package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// AllLineItemsForOrder ...
func AllLineItemsForOrder(orderID int) (lineItems []*models.LineItem, err error) {
	const lineItemsQuery = `
	SELECT pli.id, pli.quantity, pli.product_id, p.name, 
		pli.product_size_id, ps.name, ps.price, 
		pli.product_variation_id, pv.name, pv.price,
		psi.description
		FROM product_line_item AS pli
	JOIN product AS p ON pli.product_id = p.id
	JOIN product_size AS ps ON  pli.product_size_id = ps.id
	JOIN product_variation AS pv ON pli.product_variation_id = pv.id
	LEFT JOIN product_special_instruction AS psi ON pli.id = psi.product_line_item_id
	WHERE pli.order_id = ?
	`

	result, err := services.Db.DB.Query(lineItemsQuery, orderID)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var lineItem *models.LineItem = &models.LineItem{}
		var size models.ProductSize = models.ProductSize{}
		var variation models.ProductVariation = models.ProductVariation{}
		var item models.Pizza = models.Pizza{}

		err := result.Scan(&lineItem.ID, &lineItem.Quantity, &item.ID,
			&item.Name, &size.ID, &size.Name, &size.Price,
			&variation.ID, &variation.Name, &variation.Price,
			&lineItem.SpecialInstruction)

		if err != nil {
			return nil, err
		}

		lineItem.Size = &size
		lineItem.Variation = &variation
		lineItem.Item = &item

		// Add Options to it.
		lineItem.Price = size.Price + variation.Price

		lineItems = append(lineItems, lineItem)
	}

	return lineItems, nil
}
