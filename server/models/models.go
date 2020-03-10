package models

type (
	// Pizza struct
	Pizza struct {
		ID      string   `json:"ID"`
		Name    string   `json:"Name"`
		Price   float32  `json:"Price"`
		Options []Option `json:"Options"`
	}

	// Option struct
	Option struct {
		ID   string `json:"ID"`
		Name string `json:"Name"`
		Type string `json:"Type"`
	}
)
