package models

// ProductSize ...
type ProductSize struct {
	ID    string  `json:"id,omitempty"`
	Size  string  `json:"size,omitempty"`
	Price float32 `json:"price,omitempty"`
}
