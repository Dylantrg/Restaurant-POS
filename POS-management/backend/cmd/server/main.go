package main

import (
	"log"
	"net/http"

	"pos-management/internal/api"
	"pos-management/internal/config"
)

func main() {
	mux := http.NewServeMux()

	// Register API routes
	api.RegisterRoutes(mux)

	cfg := config.Load()
	addr := ":" + cfg.Port
	log.Printf("Backend server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
