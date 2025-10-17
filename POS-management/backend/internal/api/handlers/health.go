package handlers

import (
	"encoding/json"
	"net/http"
)

type healthResponse struct {
	Status string `json:"status"`
}

// HealthHandler returns basic service health.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(healthResponse{Status: "ok"})
}
