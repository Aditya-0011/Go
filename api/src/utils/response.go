package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Msg string `json:"msg"`
}


func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}