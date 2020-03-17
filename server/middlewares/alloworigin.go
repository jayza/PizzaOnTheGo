package middlewares

import (
	"net/http"
	"os"
)

// AllowOriginMiddleware struct
type AllowOriginMiddleware struct{}

// AllowOriginMiddleware ...
func (amw *AllowOriginMiddleware) AllowOriginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ACCESS_CONTROL_ALLOW_ORIGIN"))
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
