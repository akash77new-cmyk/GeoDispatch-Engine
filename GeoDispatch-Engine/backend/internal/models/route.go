package models

// Route describes the result of a routing algorithm run between two
// nodes: the path taken and its aggregate cost. Routing algorithms
// (Dijkstra, A*, implemented in later phases) will construct these.
type Route struct {
	Path          []NodeID `json:"path"`
	DistanceMeter float64  `json:"distance_meters"`
	DurationSec   float64  `json:"duration_seconds"`
	NodesExplored int      `json:"nodes_explored"`
}
