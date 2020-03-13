package models

// Order struct
type Order struct {
	ID        string     `json:"id"`
	LineItems []LineItem `json:"lineItems"`
	Status    int8       `json:"status"`
}
