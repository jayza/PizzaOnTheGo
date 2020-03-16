package models

// LineItem struct
// swagger:model
type LineItem struct {
	ID                 int               `json:"id"`
	Item               *Pizza            `json:"item"`
	UnitPrice          float64           `json:"unit_price"`
	Size               *ProductSize      `json:"size,omitempty"`
	Variation          *ProductVariation `json:"variation,omitempty"`
	Ingredients        []*Ingredient     `json:"ingredients,omitempty"`
	Quantity           int               `json:"quantity"`
	SpecialInstruction string            `json:"special_instruction,omitempty"`
}
