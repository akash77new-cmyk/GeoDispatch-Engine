// Command server is the entry point for the GeoDispatch Engine backend.
// It wires together configuration and the HTTP API and starts listening.
package main

import (
	"log"
	"net/http"

	"github.com/yourusername/geodispatch-engine/backend/internal/api"
	"github.com/yourusername/geodispatch-engine/backend/internal/config"
)

func main() {
	cfg := config.Default()

	server := api.NewServer()

	httpServer := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      server.Routes(),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	log.Printf("GeoDispatch Engine listening on :%s", cfg.Port)
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
