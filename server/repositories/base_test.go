package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
	"github.com/stretchr/testify/assert"
)

func TestAllBases(t *testing.T) {
	db := services.NewDB(models.Env{Mock: true, T: t})
	defer db.DB.Close()
	// Here we are creating rows in our mocked database.
	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(1, "Tomato sauce", 35).
		AddRow(2, "Creme fraiche", 20).
		AddRow(3, "Nothing", 60)

	services.Db.Mock.ExpectQuery(`^SELECT (.+) FROM ingredient*`).
		WillReturnRows(rows)

	bases, err := AllBases()

	if err != nil {
		t.Errorf("could not get bases: %s", err.Error())
	}

	var expectedBases []*models.Ingredient

	ingredient1 := &models.Ingredient{
		ID:    1,
		Name:  "Tomato sauce",
		Price: 35,
	}

	ingredient2 := &models.Ingredient{
		ID:    2,
		Name:  "Creme fraiche",
		Price: 20,
	}

	ingredient3 := &models.Ingredient{
		ID:    3,
		Name:  "Nothing",
		Price: 60,
	}

	expectedBases = append(expectedBases, ingredient1, ingredient2, ingredient3)

	assert.Equal(t, expectedBases, bases)
}

func TestOneBaseForPizza(t *testing.T) {
	db := services.NewDB(models.Env{Mock: true, T: t})
	defer db.DB.Close()
	// Here we are creating rows in our mocked database.
	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(1, "Tomato sauce", 35).
		AddRow(2, "Creme fraiche", 20)

	services.Db.Mock.ExpectQuery(`^SELECT (.+) FROM ingredient*`).
		WithArgs(1).
		WillReturnRows(rows)

	base, err := OneBaseForPizza(1)

	if err != nil {
		t.Errorf("could not get base: %s", err.Error())
	}

	expectedBase := &models.Ingredient{
		ID:    1,
		Name:  "Tomato sauce",
		Price: 35,
	}

	assert.Equal(t, expectedBase, base)
}
