package models

// Rider represents the party requesting a dispatch. It is intentionally
// minimal: this project is not a ride-booking app, so a rider is nothing
// more than a location a dispatch request originates from.
type Rider struct {
	ID       string   `json:"id"`
	Location Location `json:"location"`
}
