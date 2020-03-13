package controllers

import (
	"encoding/json"
	"net/http"

	repository "github.com/jayza/pizzaonthego/repositories"
)

// GetAllBasesHandler ...
func GetAllBasesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bases := repository.AllBases()
	json.NewEncoder(w).Encode(bases)
}
