package utils

import (
	"fmt"
	"net/http"
)

func IsMethodAllowed(w http.ResponseWriter, r *http.Request, allowedMethod string) bool {
	if r.Method != allowedMethod {
		message := fmt.Sprintf("only %s method is allowed", allowedMethod)
		http.Error(w, message, http.StatusMethodNotAllowed)

		return false
	}

	return true
}
