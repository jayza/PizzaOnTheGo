package helpers

import (
	"encoding/json"
	"net/http"
)

// JSONResponse ...
type JSONResponse struct {
	Data    *interface{} `json:"data,omitempty"`
	Errors  *JSONError   `json:"errors,omitempty"`
	Success bool         `json:"success"`
}

// JSONError ...
type JSONError struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

// RespondWithJSON ...
func RespondWithJSON(w http.ResponseWriter, r *http.Request, payload interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")

	var response JSONResponse

	if err == nil {
		response.Data = &payload
		response.Success = true
	} else {
		response.Success = false

		var jsonError *JSONError = &JSONError{}
		clientError, ok := err.(ClientError) // type assertion for behavior.
		if ok {
			jsonError.Code = clientError.Status()
			jsonError.Message = clientError.Error()
		} else {
			jsonError.Code = 500
			jsonError.Message = "Internal Server Error"
		}

		w.WriteHeader(jsonError.Code)

		response.Errors = jsonError
	}

	json.NewEncoder(w).Encode(response)
}
