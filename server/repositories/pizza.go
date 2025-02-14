package repositories

import (
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
)

// OnePizza ...
func OnePizza(pizzaID int) (pizza *models.Pizza, err error) {
	const pizzaQuery = `
		SELECT p.id, p.name, p.price
		FROM product as p
		WHERE p.id = ?
	`

	pizza = &models.Pizza{}

	err = services.Db.Find(pizzaQuery, services.Db.Params(pizzaID), services.Db.Fields(&pizza.ID, &pizza.Name, &pizza.Price))

	if err != nil {
		return nil, err
	}

	return pizza, nil
}

// AllPizzas ...
func AllPizzas(withIngredients bool) (pizzas []*models.Pizza, err error) {
	const pizzaQuery = `
		SELECT p.id, p.name, p.price
		FROM product as p
		WHERE product_type_id = 1
		AND custom = 0
	`

	result, err := services.Db.DB.Query(pizzaQuery)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	for result.Next() {
		var pizza *models.Pizza = &models.Pizza{}

		err := result.Scan(&pizza.ID, &pizza.Name, &pizza.Price)

		if err != nil {
			return nil, err
		}

		if withIngredients == true {
			toppings, base, dough, err := AllIngredientsForPizza(pizza.ID)

			if err != nil {
				return nil, err
			}

			pizza.Toppings = toppings
			pizza.Base = base
			pizza.Dough = dough
		}

		pizzas = append(pizzas, pizza)
	}

	return pizzas, nil
}
