package utils

import "testing"

// TestHaversineDistanceZero verifies that the distance between a point
// and itself is zero.
func TestHaversineDistanceZero(t *testing.T) {
	d := HaversineDistance(28.6139, 77.2090, 28.6139, 77.2090)
	if d != 0 {
		t.Errorf("expected 0, got %f", d)
	}
}

// TestHaversineDistanceKnown checks the distance between two well-known
// cities (New Delhi and Mumbai) against the accepted approximate value,
// allowing a tolerance since Haversine uses a spherical Earth model.
func TestHaversineDistanceKnown(t *testing.T) {
	delhiLat, delhiLng := 28.6139, 77.2090
	mumbaiLat, mumbaiLng := 19.0760, 72.8777

	d := HaversineDistance(delhiLat, delhiLng, mumbaiLat, mumbaiLng)

	const expectedKm = 1150.0
	const toleranceKm = 50.0

	gotKm := d / 1000.0
	if gotKm < expectedKm-toleranceKm || gotKm > expectedKm+toleranceKm {
		t.Errorf("expected ~%.0fkm, got %.0fkm", expectedKm, gotKm)
	}
}

// TestHaversineDistanceSymmetric verifies that distance(a, b) == distance(b, a).
func TestHaversineDistanceSymmetric(t *testing.T) {
	a := []float64{28.6139, 77.2090}
	b := []float64{19.0760, 72.8777}

	d1 := HaversineDistance(a[0], a[1], b[0], b[1])
	d2 := HaversineDistance(b[0], b[1], a[0], a[1])

	if d1 != d2 {
		t.Errorf("expected symmetric distances, got %f and %f", d1, d2)
	}
}
