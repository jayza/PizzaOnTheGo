package middlewares

import (
	"errors"
	"net/http"

	"github.com/jayza/pizzaonthego/errorshandler"
	"github.com/jayza/pizzaonthego/helpers"
)

// AuthMiddleware struct
type AuthMiddleware struct{}

// Middleware ...
// Super basic auth middleware that checks if ?loggedInAs=1 is found in the query request.
func (amw *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queryParameters := r.URL.Query()

		if queryParameters.Get("loggedInAs") == "1" {
			next.ServeHTTP(w, r)
		} else {
			e := errors.New("user does not have access to this route")
			err := errorshandler.HandleErrorCode(403, e)
			helpers.RespondWithError(w, r, 403, err)
		}
	})
}
