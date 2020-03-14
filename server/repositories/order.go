package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// OneOrder fetches one order by ID.
func OneOrder(orderID int) (order *models.Order, err error) {
	const orderQuery = `
	SELECT o.id, o.user_id, o.status FROM orders AS o
	WHERE o.id = ?
	`
	order = &models.Order{}

	err = services.Db.Find(orderQuery, services.Db.Params(&orderID), services.Db.Fields(&order.ID, &order.UserID, &order.Status))

	if err != nil {
		return nil, err
	}

	lineItems, err := AllLineItemsForOrder(orderID)

	if err != nil {
		return nil, err
	}

	order.LineItems = lineItems

	return order, nil
}

// CreateOrder ...
func CreateOrder(o models.Order) (order *models.Order, err error) {
	const orderQuery = `
	INSERT INTO orders (user_id) VALUES (?)
	`

	const lineItemQuery = `
	INSERT INTO product_line_item (order_id, product_id, product_size_id, product_variation_id, quantity) VALUES (?, ?, ?, ?, ?)
	`

	const specialInstructionQuery = `
	INSERT INTO product_special_instruction (product_line_item_id, description) VALUES (?, ?)
	`

	tx, err := services.Db.DB.Begin()

	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback()

	orderStmt, err := tx.Prepare(orderQuery)

	if err != nil {
		panic(err.Error())
	}

	defer orderStmt.Close()

	orderRes, err := orderStmt.Exec(o.UserID)

	if err != nil {
		return nil, err
	}

	lastInsertOrderID, err := orderRes.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	// LineItemQuery
	for _, lineItem := range o.LineItems {
		lineItemStmt, err := tx.Prepare(lineItemQuery)

		if err != nil {
			panic(err.Error())
		}

		defer lineItemStmt.Close()

		lineItemRes, err := lineItemStmt.Exec(lastInsertOrderID, lineItem.Item.ID, lineItem.Size.ID, lineItem.Variation.ID, lineItem.Quantity)

		if err != nil {
			return nil, err
		}

		lastInsertLineItemID, err := lineItemRes.LastInsertId()

		if err != nil {
			panic(err.Error())
		}

		// Special Instruction.
		specialInstructionStmt, err := tx.Prepare(specialInstructionQuery)

		if err != nil {
			panic(err.Error())
		}

		defer specialInstructionStmt.Close()

		_, err = specialInstructionStmt.Exec(lastInsertLineItemID, lineItem.SpecialInstruction)

		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()

	if err != nil {
		panic(err.Error())
	}
	var intLastInsertID int = int(lastInsertOrderID)

	order, err = OneOrder(intLastInsertID)

	return order, nil
}
