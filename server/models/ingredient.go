package models

// Ingredient ...
// swagger:model
type Ingredient struct {
	ID       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Category string  `json:"type,omitempty"`
	Price    float64 `json:"price,omitempty"`
}
