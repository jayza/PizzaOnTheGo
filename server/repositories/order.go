package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// OneOrder fetches one order by ID.
func OneOrder(orderID int) (order *models.Order, err error) {
	const orderQuery = `
	SELECT o.id, o.user_id, o.status, 
	osi.first_name, osi.last_name, 
	osi.phone_number, osi.street_address, 
	osi.zip_code, osi.city FROM orders AS o
	JOIN order_shipping_information AS osi ON osi.id = o.order_shipping_information_id
	WHERE o.id = ?
	`
	order = &models.Order{}

	var orderShipping *models.ShippingInfo = &models.ShippingInfo{}

	err = services.Db.Find(orderQuery, services.Db.Params(&orderID),
		services.Db.Fields(&order.ID, &order.UserID, &order.Status,
			&orderShipping.FirstName, &orderShipping.LastName,
			&orderShipping.PhoneNumber, &orderShipping.StreetAddress,
			&orderShipping.ZipCode, &orderShipping.City))

	if err != nil {
		return nil, err
	}

	order.ShippingInformation = orderShipping

	lineItems, err := AllLineItemsForOrder(orderID)

	if err != nil {
		return nil, err
	}

	order.LineItems = lineItems

	return order, nil
}

// CreateOrder ...
func CreateOrder(o models.Order) (order *models.Order, err error) {
	const orderShippingQuery = `
	INSERT INTO order_shipping_information (first_name, last_name, phone_number, street_address, zip_code, city) VALUES (?, ?, ?, ?, ?, ?)
	`
	const orderQuery = `
	INSERT INTO orders (user_id, order_shipping_information_id) VALUES (?, ?)
	`

	const lineItemQuery = `
	INSERT INTO product_line_item (order_id, product_id, product_size_id, product_variation_id, quantity, unit_price) VALUES (?, ?, ?, ?, ?, ?)
	`

	const specialInstructionQuery = `
	INSERT INTO product_special_instruction (product_line_item_id, description) VALUES (?, ?)
	`

	tx, err := services.Db.DB.Begin()

	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	orderShippingStmt, err := tx.Prepare(orderShippingQuery)

	if err != nil {
		return nil, err
	}

	defer orderShippingStmt.Close()

	orderShippingRes, err := orderShippingStmt.Exec(o.ShippingInformation.FirstName,
		o.ShippingInformation.LastName, o.ShippingInformation.PhoneNumber,
		o.ShippingInformation.StreetAddress, o.ShippingInformation.ZipCode, o.ShippingInformation.City)

	if err != nil {
		return nil, err
	}

	lastInsertOrderShippingID, err := orderShippingRes.LastInsertId()

	orderStmt, err := tx.Prepare(orderQuery)

	if err != nil {
		return nil, err
	}

	defer orderStmt.Close()

	orderRes, err := orderStmt.Exec(o.UserID, lastInsertOrderShippingID)

	if err != nil {
		return nil, err
	}

	lastInsertOrderID, err := orderRes.LastInsertId()

	if err != nil {
		return nil, err
	}

	// LineItemQuery
	for _, lineItem := range o.LineItems {
		var args []interface{}

		args = append(args, lineItem.Size.ID, lineItem.Variation.ID, lineItem.Item.ID)

		getPriceSumQuery := "SELECT SUM(prices.price) AS unit_price FROM " +
			"(SELECT price FROM product_size WHERE product_size.id = ? " +
			"UNION ALL " +
			"SELECT price FROM product_variation WHERE product_variation.id = ? " +
			"UNION ALL " +
			"SELECT price FROM product WHERE product.id = ? "

		// Extra Options parameter
		if len(lineItem.Ingredients) > 0 {
			getPriceSumQuery = getPriceSumQuery +
				"UNION ALL " +
				"SELECT price FROM ingredient " +
				"WHERE ingredient.id IN (?" + strings.Repeat(",?", len(lineItem.Ingredients)-1) + ")"

			var opts []interface{}
			if len(lineItem.Ingredients) > 0 {
				for _, opt := range lineItem.Ingredients {
					opts = append(opts, opt.ID)
				}
			}

			args = append(args, opts...)
		}

		getPriceSumQuery = getPriceSumQuery + ") prices"

		// Get Prices from variation, size, options
		var lineItemPrice float64 = 0

		if priceRes := tx.QueryRow(getPriceSumQuery, args...).Scan(&lineItemPrice); priceRes == sql.ErrNoRows {
			return nil, err
		}

		// LineItem
		lineItemStmt, err := tx.Prepare(lineItemQuery)

		if err != nil {
			return nil, err
		}

		defer lineItemStmt.Close()

		lineItemRes, err := lineItemStmt.Exec(lastInsertOrderID, lineItem.Item.ID, lineItem.Size.ID, lineItem.Variation.ID, lineItem.Quantity, lineItemPrice)

		if err != nil {
			return nil, err
		}

		lastInsertLineItemID, err := lineItemRes.LastInsertId()

		if err != nil {
			return nil, err
		}

		// Extra Ingredients
		if len(lineItem.Ingredients) > 0 {
			extraIngredientsQuery := "INSERT INTO product_extra_ingredients (product_line_item_id, ingredient_id) VALUES (?,?)" + strings.Repeat(" ,(?,?)", len(lineItem.Ingredients)-1)
			extraIngredientsStmt, err := tx.Prepare(extraIngredientsQuery)

			fmt.Println(extraIngredientsQuery)
			if err != nil {
				return nil, err
			}

			var args []interface{}

			for _, ingredient := range lineItem.Ingredients {
				args = append(args, lastInsertLineItemID, ingredient.ID)
			}

			fmt.Println(args)

			defer extraIngredientsStmt.Close()

			_, err = extraIngredientsStmt.Exec(args...)

			if err != nil {
				return nil, err
			}
		}

		// Special Instruction.
		specialInstructionStmt, err := tx.Prepare(specialInstructionQuery)

		if err != nil {
			return nil, err
		}

		defer specialInstructionStmt.Close()

		_, err = specialInstructionStmt.Exec(lastInsertLineItemID, lineItem.SpecialInstruction)

		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	var intLastInsertID int = int(lastInsertOrderID)

	order, err = OneOrder(intLastInsertID)

	return order, nil
}
