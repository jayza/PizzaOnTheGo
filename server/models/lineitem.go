package models

// LineItem struct
type LineItem struct {
	ID       string      `json:"id"`
	Product  Pizza       `json:"product"`
	Options  PizzaOption `json:"options"`
	Quantity uint8       `json:"quantity"`
}
