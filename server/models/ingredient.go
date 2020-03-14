package models

// Ingredient ...
type Ingredient struct {
	ID       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Category string  `json:"type,omitempty"`
	Price    float32 `json:"price,omitempty"`
}
