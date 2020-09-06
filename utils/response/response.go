package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

// Success --> Returns successfull response
func Success(w http.ResponseWriter, statusCode int, response interface{}) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Error --> Returns error response
func Error(w http.ResponseWriter, statusCode int, err interface{}) {
	body, _ := json.Marshal(errorResponse{Message: fmt.Sprint(err)})
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}
