package utils

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, status int, message string, error error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := map[string]interface{}{
		"error": message + ": " + error.Error(),
		"data":  nil,
	}
	json.NewEncoder(w).Encode(resp)
}

func SuccessResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := map[string]interface{}{
		"success": message,
		"data":    data,
	}
	json.NewEncoder(w).Encode(resp)
}
