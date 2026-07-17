package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yourusername/geodispatch-engine/backend/internal/api"
)

// TestHealthEndpoint verifies GET /health returns HTTP 200 with a
// well-formed status payload. This is an integration test: it exercises
// the real router rather than calling the handler function directly.
func TestHealthEndpoint(t *testing.T) {
	server := api.NewServer()
	ts := httptest.NewServer(server.Routes())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var body struct {
		Status    string  `json:"status"`
		UptimeSec float64 `json:"uptime_seconds"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body.Status != "ok" {
		t.Errorf("expected status 'ok', got %q", body.Status)
	}
	if body.UptimeSec < 0 {
		t.Errorf("expected non-negative uptime, got %f", body.UptimeSec)
	}
}
