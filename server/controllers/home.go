package controllers

import (
	"fmt"
	"log"
	"net/http"
)

/*
 Miscellaneous Handlers
*/

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

/*
Topping Route Handler
*/

// GetAllToppingsHandler ...
// func GetAllToppingsHandler(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(toppings)
// }
