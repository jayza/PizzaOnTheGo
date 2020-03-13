package controllers

import (
	"encoding/json"
	"net/http"

	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllDoughsHandler ...
func GetAllDoughsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	doughs := repository.AllDoughs()
	json.NewEncoder(w).Encode(doughs)
}
