package api

import (
	"encoding/json"
	"net/http"
)

// writeJSON encodes v as JSON and writes it to w with the given status
// code. Centralizing this avoids repeating header/encoding boilerplate
// in every handler.
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// errorResponse is the standard JSON shape for error replies.
type errorResponse struct {
	Error string `json:"error"`
}

// writeError writes a JSON error response.
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Error: message})
}
