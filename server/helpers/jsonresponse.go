package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/jayza/pizzaonthego/errors"
)

// JSONResponse ...
type JSONResponse struct {
	Data   *interface{} `json:"data,omitempty"`
	Errors *JSONError   `json:"errors,omitempty"`
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
	} else {
		err = errors.HandleError(err)

		var jsonError *JSONError = &JSONError{}

		clientError, ok := err.(errors.ClientError) // type assertion for behavior.
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

// RespondWithError ...
func RespondWithError(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")

	var response JSONResponse

	err = errors.HandleError(err)

	var jsonError *JSONError = &JSONError{}

	clientError, ok := err.(errors.ClientError) // type assertion for behavior.
	if ok {
		jsonError.Code = clientError.Status()
		jsonError.Message = clientError.Error()
	} else {
		jsonError.Code = 500
		jsonError.Message = "Internal Server Error"
	}

	w.WriteHeader(jsonError.Code)

	response.Errors = jsonError

	json.NewEncoder(w).Encode(response)
}
