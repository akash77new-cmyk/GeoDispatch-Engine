# Design Notes

This document tracks the reasoning behind non-obvious design decisions
so they can be defended in an interview setting. It is updated as each
phase adds new decisions.

## Why no external algorithm libraries

KD-Tree, Geohashing, Dijkstra, A*, Haversine distance, and K-nearest
search are all implemented from scratch. The point of this project is to
demonstrate data structures and algorithms knowledge directly — pulling
in a library would hide the exact thing the project exists to show.

## Why no database

All state (drivers, graph, routes) lives in memory, guarded by
`sync.RWMutex` where concurrent access is possible. This keeps the
project's scope on algorithms and concurrency rather than persistence
and schema design, and it matches how a hot-path spatial index actually
behaves in production systems (in-memory with periodic snapshots), even
though this project doesn't implement the snapshotting.

## Why the spatial search strategies share one interface

See `docs/architecture.md#why-an-interface-for-spatial-search`. In short:
it lets Brute Force, Geohash, and KD-Tree be benchmarked apples-to-apples
and swapped without touching dispatch logic — the same pattern real
systems use to A/B test index implementations in production.

## Why Haversine distance is used twice

1. As the distance metric inside spatial search (comparing straight-line
   proximity between rider and drivers).
2. As the A* heuristic once routing is implemented (Phase 8).

For (2), Haversine distance is **admissible**: since it is a great-circle
distance, it can never overestimate the true road-network distance
between two points (roads are never shorter than a straight line). An
admissible heuristic is what guarantees A* still finds the optimal path,
while typically expanding far fewer nodes than Dijkstra because it is
guided toward the goal instead of expanding uniformly in all directions.
This tradeoff, and measured node-expansion counts, are documented further
in `docs/benchmarks.md` once Phase 8 and Phase 12 land.

## Why `Driver.Status` is an enum, not a bool

An `is_available bool` field can only ever represent two states. Modeling
status as a `DriverStatus` string type leaves room to add `en_route` or
`on_break` later without a breaking schema change to the model or API.

## Why the frontend has no state management library

The frontend is a visualization layer, not the product. Its state (driver
markers, selected driver, route polyline) is shallow and local to the
Dispatch Simulator page, so React's built-in `useState`/`useEffect` is
sufficient; adding Redux or similar would be unjustified complexity for
what this frontend actually does.

## Concurrency strategy (introduced fully in Phase 11)

The spatial index and driver registry will be protected by
`sync.RWMutex`: reads (K-nearest queries, which vastly outnumber writes
in a dispatch system) take a read lock and can proceed concurrently with
each other; writes (driver insert/update/remove) take a write lock.
Double-assignment of a driver during concurrent dispatch requests is
prevented by treating "mark driver as assigned" as part of the same
critical section as "select driver" — a driver is checked and claimed
atomically, not checked-then-claimed as two separate steps, which would
be a classic check-then-act race condition.
