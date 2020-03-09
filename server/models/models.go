package models

type (
	// Pizza struct
	Pizza struct {
		ID    string `json:"ID"`
		Name  string `json:"Name"`
		Price uint32 `json:"Price"`
	}

	// Topping struct
	Topping struct {
		ID   string `json:"ID"`
		Name string `json:"Name"`
	}
)
