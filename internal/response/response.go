package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status   string      `json:"status"`
	Messages string      `json:"messages"`
	Data     interface{} `json:"data,omitempty"`
	Error    interface{} `json:"error,omitempty"`
}

// Response sukses
func Success(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Status:   "success",
		Messages: message,
		Data:     data,
	})
}

// Response error
func Error(w http.ResponseWriter, statusCode int, message string, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		Status:   "error",
		Messages: message,
		Error:    err,
	})
}
