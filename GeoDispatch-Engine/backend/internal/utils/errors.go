package utils

import "errors"

// Sentinel errors shared across packages. Using sentinel errors (rather
// than ad-hoc string errors) lets callers use errors.Is for reliable
// error-type checks instead of brittle string matching.
var (
	ErrNotFound      = errors.New("resource not found")
	ErrInvalidInput  = errors.New("invalid input")
	ErrAlreadyExists = errors.New("resource already exists")
	ErrNoPath        = errors.New("no path exists between nodes")
	ErrNoDrivers     = errors.New("no drivers available")
)
