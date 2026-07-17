// Package config centralizes all runtime configuration for the GeoDispatch
// Engine backend. Keeping configuration in one place (rather than scattering
// magic numbers across packages) makes the system's tunable parameters
// explicit and easy to reason about during code review.
package config

import "time"

// Config holds every tunable parameter used by the server and its
// subsystems. Fields are grouped by the subsystem that consumes them.
type Config struct {
	// Server
	Port            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration

	// Dispatch
	// DefaultCandidateCount is the "K" used when shortlisting nearby
	// drivers via the spatial index before running road-network routing.
	DefaultCandidateCount int

	// MaxSearchRadiusMeters bounds how far the spatial index will look
	// for candidate drivers around a rider before giving up.
	MaxSearchRadiusMeters float64

	// Geohash
	// GeohashPrecision controls the length of the geohash string used
	// for bucketing drivers. Higher precision => smaller buckets.
	GeohashPrecision int

	// Simulation
	DefaultDriverCount int
}

// Default returns the configuration used when the server starts without
// any environment overrides. Centralizing defaults here means every
// subsystem can be tuned from a single source of truth during development.
func Default() *Config {
	return &Config{
		Port:            "8080",
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		ShutdownTimeout: 5 * time.Second,

		DefaultCandidateCount: 5,
		MaxSearchRadiusMeters: 5000,

		GeohashPrecision: 6,

		DefaultDriverCount: 1000,
	}
}
