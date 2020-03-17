package models

// ProductSize ...
// swagger:model
type ProductSize struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}
