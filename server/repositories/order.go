package repositories

import (
	"fmt"

	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// OneOrder fetches one order by ID.
func OneOrder(orderID int) (order *models.Order, err error) {
	const orderQuery = `
	SELECT o.id, o.user_id, o.status
	FROM orders as o
	WHERE o.id = ?
	`
	order = &models.Order{}

	err = services.Db.Find(orderQuery, services.Db.Params(&orderID), services.Db.Fields(&order.ID, &order.UserID, &order.Status))

	if err != nil {
		return nil, err
	}

	return order, nil
}

// CreateOrder ...
func CreateOrder(o models.Order) (order *models.Order, err error) {
	const orderQuery = `
	INSERT INTO orders (user_id) VALUES (?)
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

	lastInsertID, err := orderRes.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	err = tx.Commit()

	if err != nil {
		panic(err.Error())
	}

	var intLastInsertID int = int(lastInsertID)
	fmt.Println("Last ID:", lastInsertID)
	fmt.Println("Int Last ID:", intLastInsertID)

	order, err = OneOrder(intLastInsertID)

	return order, nil
}
