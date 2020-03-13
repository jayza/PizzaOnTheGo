package controllers

import (
	"encoding/json"
	"net/http"

	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllCrustsHandler ...
func GetAllCrustsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	crusts := repository.AllCrusts()
	json.NewEncoder(w).Encode(crusts)
}
