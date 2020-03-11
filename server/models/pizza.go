package models

// Pizza struct
type Pizza struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Price    float32       `json:"price"`
	Toppings []PizzaOption `json:"toppings"`
	Crust    PizzaOption   `json:"crust"`
	Dough    PizzaOption   `json:"dough"`
	Base     PizzaOption   `json:"base"`
	Size     PizzaOption   `json:"size"`
}
