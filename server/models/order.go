package models

// Order struct
// swagger:model
type Order struct {
	ID                  int           `json:"id,omitempty"`
	LineItems           []*LineItem   `json:"lineItems,omitempty"`
	Status              *int          `json:"status,omitempty"`
	UserID              int           `json:"userId,omitempty"`
	ShippingInformation *ShippingInfo `json:"shippingInformation,omitempty"`
}

// OrderCreateParams ...
// swagger:parameters addOrder
type OrderCreateParams struct {
	LineItems           []*LineItem   `json:"lineItems,omitempty"`
	UserID              int           `json:"userId,omitempty"`
	ShippingInformation *ShippingInfo `json:"shippingInformation,omitempty"`
}

// OrderFindParams ...
// swagger:parameters findOrder
type OrderFindParams struct {
	ID int `json:"id,omitempty"`
}
