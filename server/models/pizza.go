package models

// Pizza struct
type Pizza struct {
	ID       string            `json:"id,omitempty"`
	Name     string            `json:"name,omitempty"`
	Price    float32           `json:"price,omitempty"`
	Toppings []*Ingredient     `json:"toppings,omitempty"`
	Crust    *ProductVariation `json:"crust,omitempty"`
	Dough    *Ingredient       `json:"dough,omitempty"`
	Base     *Ingredient       `json:"base,omitempty"`
	Size     *ProductSize      `json:"size,omitempty"`
}
