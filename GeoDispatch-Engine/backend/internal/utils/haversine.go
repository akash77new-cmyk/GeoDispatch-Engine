package utils

import "math"

// earthRadiusMeters is the mean radius of the Earth, used by the
// Haversine formula. A mean radius is accurate enough for routing and
// dispatch use cases; it is not survey-grade geodesy.
const earthRadiusMeters = 6371000.0

// HaversineDistance returns the great-circle distance, in meters,
// between two WGS84 coordinates.
//
// It is used in two distinct roles in this system:
//  1. As the distance metric for spatial search (KD-Tree, brute force).
//  2. As the admissible heuristic for A* search (Phase 8), since the
//     straight-line distance between two points can never exceed the
//     true road-network distance between them.
func HaversineDistance(lat1, lng1, lat2, lng2 float64) float64 {
	phi1 := degToRad(lat1)
	phi2 := degToRad(lat2)
	deltaPhi := degToRad(lat2 - lat1)
	deltaLambda := degToRad(lng2 - lng1)

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
		math.Cos(phi1)*math.Cos(phi2)*
			math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusMeters * c
}

func degToRad(deg float64) float64 {
	return deg * math.Pi / 180.0
}
