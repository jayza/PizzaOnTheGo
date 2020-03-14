package models

// ProductVariation ...
type ProductVariation struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"size,omitempty"`
	Price float64 `json:"price,omitempty"`
}
