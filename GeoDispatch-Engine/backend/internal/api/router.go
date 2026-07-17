// Package api wires up HTTP handlers for the GeoDispatch Engine. Routing
// is kept on the standard library's net/http.ServeMux (Go 1.22+ supports
// method-aware patterns natively) so the project has zero external web
// framework dependencies.
package api

import (
	"net/http"
	"time"
)

// Server bundles the dependencies HTTP handlers need. As dispatch,
// spatial, and routing engines are implemented in later phases, they
// will be added here as interface-typed fields and injected in main.go.
// Keeping dependencies as an explicit struct (rather than package-level
// globals) keeps handlers testable in isolation.
type Server struct {
	startedAt time.Time
}

// NewServer constructs a Server with its dependencies.
func NewServer() *Server {
	return &Server{startedAt: time.Now()}
}

// Routes returns the fully configured HTTP handler for the application.
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", s.handleHealth)

	// The following endpoints are part of the target API surface but are
	// intentionally not yet implemented; they will be filled in as their
	// owning subsystems are built in later phases:
	//   POST   /drivers
	//   PUT    /drivers/{id}/location
	//   DELETE /drivers/{id}
	//   POST   /dispatch
	//   GET    /benchmark

	return withLogging(mux)
}

// healthResponse is the JSON payload returned by GET /health.
type healthResponse struct {
	Status    string  `json:"status"`
	UptimeSec float64 `json:"uptime_seconds"`
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	resp := healthResponse{
		Status:    "ok",
		UptimeSec: time.Since(s.startedAt).Seconds(),
	}
	writeJSON(w, http.StatusOK, resp)
}
