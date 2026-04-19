package utils

import (
	"encoding/json"
	"net/http"
)

func JSON(responseWriter http.ResponseWriter, data any, statusCode int) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	json.NewEncoder(responseWriter).Encode(data)
}

func Error(responseWriter http.ResponseWriter, message string, statusCode int) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	json.NewEncoder(responseWriter).Encode(map[string]string{"error": message})
}