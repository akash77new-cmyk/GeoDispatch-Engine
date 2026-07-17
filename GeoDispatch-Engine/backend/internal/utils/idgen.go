package utils

import (
	"crypto/rand"
	"fmt"
)

// NewID generates a short random hex identifier, used for drivers and
// riders created without a caller-supplied ID (e.g. by the simulator).
func NewID(prefix string) string {
	b := make([]byte, 6)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%s-%x", prefix, b)
}
