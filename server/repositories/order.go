package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// OneOrder fetches one order by ID.
func OneOrder(orderID string) (order *models.Order, err error) {
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

// UpsertOrder ...
func UpsertOrder() {

}
