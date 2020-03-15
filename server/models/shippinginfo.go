package models

// ShippingInfo struct
// swagger:model
type ShippingInfo struct {
	ID            int    `json:"id,omitempty"`
	FirstName     string `json:"firstName,omitempty"`
	LastName      string `json:"lastName,omitempty"`
	PhoneNumber   string `json:"phone,omitempty"`
	StreetAddress string `json:"streetAddress,omitempty"`
	ZipCode       string `json:"zipCode,omitempty"`
	City          string `json:"city,omitempty"`
}
