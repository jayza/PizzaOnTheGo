package models

// ProductVariation ...
// swagger:model
type ProductVariation struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"size,omitempty"`
	Price float64 `json:"price,omitempty"`
}
