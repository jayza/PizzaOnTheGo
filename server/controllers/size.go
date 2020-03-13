package controllers

import (
	"encoding/json"
	"net/http"

	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllSizesHandler ...
func GetAllSizesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sizes := repository.AllSizes()
	json.NewEncoder(w).Encode(sizes)
}
