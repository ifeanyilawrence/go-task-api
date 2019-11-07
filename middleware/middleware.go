package middleware

import (
	"net/http"

	"github.com/ifeanyilawrence/go-task-api/controllers"
)

//AuthMiddleware : authorization middleware, can be refactored out
func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if status, _ := controllers.Authenticate(r); !status {
			http.Error(w, http.StatusText(401), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}
