package api

import (
	"log"
	"net/http"
	"time"
)

// withLogging wraps a handler with basic structured request logging.
// This is deliberately minimal for Phase 2; it exists so every future
// endpoint gets consistent observability for free.
func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
