package errorshandler

import (
	"database/sql"
)

// ClientError is an error whose details to be shared with client.
type ClientError interface {
	Error() string
	Status() int
}

// HTTPError implements ClientError interface.
type HTTPError struct {
	Cause   error  `json:"-"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

// Status ...
func (e *HTTPError) Status() int {
	return e.Code
}

// NewHTTPError ...
func NewHTTPError(err error, code int, msg string) error {
	return &HTTPError{
		Cause:   err,
		Message: msg,
		Code:    code,
	}
}

// HandleError takes an error parameter and determines which kind of HTTPError Response to return.
func HandleError(err error) error {
	if _, ok := err.(*HTTPError); ok {
		return err
	}

	switch err {
	case sql.ErrNoRows:
		return HandleErrorCode(404, err)
	default:
		return HandleErrorCode(500, err)
	}
}

// HandleErrorCode ...
func HandleErrorCode(code int, err error) error {
	switch code {
	case 400:
		return NewHTTPError(err, 400, "Bad Request")
	case 403:
		return NewHTTPError(err, 403, "Forbidden")
	case 404:
		return NewHTTPError(err, 404, "Not Found")
	case 405:
		return NewHTTPError(err, 405, "Method Not Allowed")
	case 500:
		return NewHTTPError(err, 500, "Internal Server Error")
	default:
		return NewHTTPError(err, 500, "Internal Server Error")
	}
}
