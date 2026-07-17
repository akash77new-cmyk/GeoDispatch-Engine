package models

// Location represents a point on Earth's surface using decimal-degree
// WGS84 coordinates. It is the fundamental geospatial primitive shared
// by drivers, riders, and graph nodes throughout the system.
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
