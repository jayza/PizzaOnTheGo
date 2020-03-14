package helpers

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
