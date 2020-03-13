package models

// Ingredient ...
type Ingredient struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"type"`
	Price    float32 `json:"price"`
}
