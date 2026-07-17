package models

import "time"

// DriverStatus describes the current availability state of a driver.
// Modeling this as an explicit enum (rather than a bool) leaves room to
// add states like "en_route" or "on_break" without breaking callers.
type DriverStatus string

const (
	DriverAvailable DriverStatus = "available"
	DriverAssigned  DriverStatus = "assigned"
	DriverOffline   DriverStatus = "offline"
)

// Driver represents a vehicle available for dispatch. Location is updated
// frequently (e.g. on a GPS ping), so spatial indexes must support cheap
// updates in addition to inserts and removals.
type Driver struct {
	ID        string       `json:"id"`
	Location  Location     `json:"location"`
	Status    DriverStatus `json:"status"`
	UpdatedAt time.Time    `json:"updated_at"`
}
