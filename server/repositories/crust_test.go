package repositories

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jayza/pizzaonthego/models"
	"github.com/jayza/pizzaonthego/services"
	"github.com/stretchr/testify/assert"
)

func TestAllCrusts(t *testing.T) {
	db := services.NewDB(models.Env{Mock: true, T: t})
	defer db.DB.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(1, "Super cheesy", 35).
		AddRow(2, "Thick", 20).
		AddRow(3, "Thin", 60)

	services.Db.Mock.ExpectQuery(`^SELECT (.+) FROM product_variation*`).
		WillReturnRows(rows)

	crusts, err := AllCrusts()

	if err != nil {
		t.Errorf("could not get crusts: %s", err.Error())
	}

	var expectedCrusts []*models.ProductVariation

	variation1 := &models.ProductVariation{
		ID:    1,
		Name:  "Super cheesy",
		Price: 35,
	}

	variation2 := &models.ProductVariation{
		ID:    2,
		Name:  "Thick",
		Price: 20,
	}

	variation3 := &models.ProductVariation{
		ID:    3,
		Name:  "Thin",
		Price: 60,
	}

	expectedCrusts = append(expectedCrusts, variation1, variation2, variation3)

	assert.Equal(t, expectedCrusts, crusts)
}

func TestOneCrust(t *testing.T) {
	db := services.NewDB(models.Env{Mock: true, T: t})
	defer db.DB.Close()
	// Here we are creating rows in our mocked database.
	rows := sqlmock.NewRows([]string{"id", "name", "price"}).
		AddRow(1, "Thick", 35).
		AddRow(2, "Thin", 20)

	services.Db.Mock.ExpectQuery(`^SELECT (.+) FROM product_variation*`).
		WithArgs(1).
		WillReturnRows(rows)

	crust, err := OneCrust(1)

	if err != nil {
		t.Errorf("could not get base: %s", err.Error())
	}

	expectedCrust := &models.ProductVariation{
		ID:    2,
		Name:  "Thin",
		Price: 20,
	}

	assert.Equal(t, expectedCrust, crust)
}
