package models

// Pizza struct
type Pizza struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	Price    float32          `json:"price"`
	Toppings []Ingredient     `json:"toppings"`
	Crust    ProductVariation `json:"crust"`
	Dough    Ingredient       `json:"dough"`
	Base     Ingredient       `json:"base"`
	Size     ProductSize      `json:"size"`
}
