package pizza

import (
	"github.com/jayza/pizzaonthego/models"
	database "github.com/jayza/pizzaonthego/services"
)

// GetOne ...
func GetOne(id string) models.Pizza {
	var db = database.Connect()

	result, err := db.Query("SELECT ID, Name, Price FROM Pizza WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var pizza models.Pizza

	for result.Next() {
		err := result.Scan(&pizza.ID, &pizza.Name, &pizza.Price)
		if err != nil {
			panic(err.Error())
		}
	}

	db.Close()

	return pizza
}

// GetAll ...
func GetAll() []models.Pizza {
	var db = database.Connect()

	result, err := db.Query("SELECT ID, Name, Price FROM Pizza")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var pizzas []models.Pizza

	for result.Next() {
		var pizza models.Pizza
		err := result.Scan(&pizza.ID, &pizza.Name, &pizza.Price)
		if err != nil {
			panic(err.Error())
		}
		pizzas = append(pizzas, pizza)
	}

	db.Close()

	return pizzas
}
