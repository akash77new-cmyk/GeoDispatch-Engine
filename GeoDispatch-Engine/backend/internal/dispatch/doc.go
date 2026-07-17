// Package dispatch will implement the dispatch engine that orchestrates
// spatial search and routing to select the best driver for a rider
// request. Implemented in Phase 9.
//
// The dispatch engine will depend on the spatial.Index interface (not a
// concrete implementation), so Brute Force, Geohash, or KD-Tree search
// can be swapped without changing dispatch logic.
package dispatch
