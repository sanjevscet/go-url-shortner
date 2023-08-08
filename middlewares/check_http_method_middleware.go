package middlewares

import (
	"fmt"
	"net/http"
)

func IsHttpMethodAllowed(allowedMethod string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowedMethod {
			errorMsg := fmt.Sprintf("only %s method is allowed", allowedMethod)
			http.Error(w, errorMsg, http.StatusBadRequest)

			return
		}

		next(w, r)
	}
}
