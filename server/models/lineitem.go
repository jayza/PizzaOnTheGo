package models

// LineItem struct
type LineItem struct {
	ID                 int               `json:"id"`
	Item               *Pizza            `json:"item"`
	Price              float64           `json:"price"`
	Size               *ProductSize      `json:"size,omitempty"`
	Variation          *ProductVariation `json:"variation,omitempty"`
	Quantity           uint8             `json:"quantity"`
	SpecialInstruction string            `json:"special_instruction,omitempty"`
}
