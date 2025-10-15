package api

import (
	"net/http"
	"pos-management/internal/api/handlers"
	"pos-management/internal/api/middleware"
)

// RegisterRoutes wires up all API routes and middleware.
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/health", middleware.Chain(handlers.HealthHandler))
	// Root landing page
	mux.HandleFunc("/", middleware.Chain(handlers.RootHandler))
}
