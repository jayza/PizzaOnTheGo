package models

// ProductVariation ...
type ProductVariation struct {
	ID    string  `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float32 `json:"price,omitempty"`
}
