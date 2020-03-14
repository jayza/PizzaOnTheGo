package models

// Order struct
type Order struct {
	ID        string      `json:"id,omitempty"`
	LineItems []*LineItem `json:"lineItems,omitempty"`
	Status    int8        `json:"status,omitempty"`
	UserID    string      `json:"userId,omitempty"`
}
