package models

// Order struct
// swagger:model
type Order struct {
	ID                  int           `json:"id,omitempty"`
	LineItems           []*LineItem   `json:"lineItems,omitempty"`
	Status              *int          `json:"status,omitempty"`
	UserID              string        `json:"userId,omitempty"`
	ShippingInformation *ShippingInfo `json:"shippingInformation,omitempty"`
}
