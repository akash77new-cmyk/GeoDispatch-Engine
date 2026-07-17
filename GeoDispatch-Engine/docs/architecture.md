# Architecture

## Overview

GeoDispatch Engine is split into two independently runnable pieces:

- **Backend** (`/backend`) — a Go service that owns all data structures,
  algorithms, and state. This is where the engineering content of the
  project lives.
- **Frontend** (`/frontend`) — a lightweight React + Leaflet visualization
  layer. It renders what the backend computes; it does not contain
  business logic of its own.

## Backend Layering

```
cmd/server        entry point: wiring only, no logic
internal/api       HTTP handlers, routing, request/response shaping
internal/dispatch  orchestrates spatial search + routing -> best driver
internal/spatial   pluggable nearest-driver search (brute force / geohash / kdtree)
internal/routing   Dijkstra and A* over the road network
internal/graph     road-network graph construction from OSM data
internal/models    shared data types (Driver, Rider, Graph, Route, ...)
internal/simulation synthetic driver/rider generation for demos
internal/benchmark  timing/quality comparison harness
internal/config    centralized tunable parameters
internal/utils     cross-cutting helpers (Haversine, IDs, error types)
```

Dependencies point inward: `api` depends on `dispatch`, `dispatch` depends
on `spatial` and `routing` *interfaces*, and concrete spatial/routing
implementations depend on `models`. Nothing in `models` or `utils`
depends on anything else in the project — they are the shared vocabulary
every other package builds on.

## Why an Interface for Spatial Search

`internal/spatial` defines a single `Index` interface:

```go
type Index interface {
    Insert(d models.Driver) error
    Remove(id string) error
    Update(id string, loc models.Location) error
    KNearest(loc models.Location, k int) ([]models.Driver, error)
}
```

`internal/dispatch` is written against this interface only. That means:

- Brute Force (Phase 3), Geohash (Phase 4), and KD-Tree (Phase 5) are all
  drop-in replacements for each other.
- The benchmark suite (Phase 12) can run the exact same dispatch logic
  against all three implementations and compare them fairly.
- Adding a fourth spatial strategy later (e.g. an R-tree) requires no
  changes to dispatch or API code.

## Dispatch Pipeline

```
Rider Request
     |
     v
Spatial Index  (Insert/Remove/Update happen concurrently, protected by
     |          a sync.RWMutex — see docs/design.md)
     v
Top-K Candidate Drivers
     |
     v
A* Routing (per candidate, using Haversine as the heuristic)
     |
     v
Compare ETA, pick minimum
     |
     v
Dispatch Response
```

## Request Flow (current, Phase 2)

```
Browser -> Vite dev server (proxies /api/*) -> Go net/http.ServeMux
                                                     |
                                                     v
                                              api.Server handlers
                                                     |
                                                     v
                                            GET /health (implemented)
```

Later phases extend this same flow to `/drivers`, `/dispatch`, and
`/benchmark` without changing how the frontend talks to the backend.
