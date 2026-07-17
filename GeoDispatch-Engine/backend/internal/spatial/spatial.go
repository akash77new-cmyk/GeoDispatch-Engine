// Package spatial defines the shared contract that every spatial search
// strategy (brute force, geohash, KD-tree) implements. The dispatch
// engine depends only on this interface, so search strategies are
// interchangeable without touching dispatch logic.
package spatial

import "github.com/yourusername/geodispatch-engine/backend/internal/models"

// Index is the interface every spatial index implementation must
// satisfy. Concrete implementations live in the bruteforce, geohash, and
// kdtree subpackages, added in Phases 3-5.
type Index interface {
	// Insert adds a driver to the index.
	Insert(d models.Driver) error

	// Remove deletes a driver from the index by ID.
	Remove(id string) error

	// Update changes a driver's indexed location.
	Update(id string, loc models.Location) error

	// KNearest returns the k drivers closest to loc.
	KNearest(loc models.Location, k int) ([]models.Driver, error)
}
