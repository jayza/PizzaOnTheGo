package models

// Pizza struct
// swagger:model
type Pizza struct {
	ID       int               `json:"id,omitempty"`
	Name     string            `json:"name,omitempty"`
	Price    float64           `json:"price,omitempty"`
	Toppings []*Ingredient     `json:"toppings,omitempty"`
	Crust    *ProductVariation `json:"crust,omitempty"`
	Dough    *Ingredient       `json:"dough,omitempty"`
	Base     *Ingredient       `json:"base,omitempty"`
	Size     *ProductSize      `json:"size,omitempty"`
}

// FindPizzaParams struct
// swagger:parameters findPizza findPizzaToppings
type FindPizzaParams struct {
	ID int `json:"id,omitempty"`
}
