package models

// DispatchRequest is the payload for POST /dispatch: a rider location
// plus optional tuning parameters for the search.
type DispatchRequest struct {
	Rider          Location `json:"rider"`
	CandidateCount int      `json:"candidate_count,omitempty"`
}

// DispatchResult is the response for a successful dispatch: the chosen
// driver plus enough metadata to explain and visualize the decision.
type DispatchResult struct {
	SelectedDriver   *Driver  `json:"selected_driver"`
	ETASeconds       float64  `json:"eta_seconds"`
	DistanceMeters   float64  `json:"distance_meters"`
	Route            []NodeID `json:"route,omitempty"`
	AlgorithmUsed    string   `json:"algorithm_used"`
	CandidateDrivers []Driver `json:"candidate_drivers"`
	DispatchTimeMs   float64  `json:"dispatch_time_ms"`
}
