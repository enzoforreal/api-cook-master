package handlers

import (
	"encoding/json"
	"net/http"
)

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	errorMessage := map[string]string{"error": "Method not allowed"}
	json.NewEncoder(w).Encode(errorMessage)
}
