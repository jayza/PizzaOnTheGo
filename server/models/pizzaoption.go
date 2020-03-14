package models

// PizzaOption struct
type PizzaOption struct {
	ID    string  `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Type  string  `json:"type,omitempty"`
	Price float32 `json:"price,omitempty"`
}
