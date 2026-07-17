# GeoDispatch Engine

**Geospatial driver discovery and ETA-optimized dispatch, built from scratch.**

GeoDispatch Engine is a backend systems project that reimplements the core
geospatial matching problem at the heart of ride-hailing platforms: given a
rider and thousands of drivers, find the driver who can arrive fastest —
using hand-built spatial indexes and graph search rather than off-the-shelf
libraries.

This is **not** a ride-booking app. There is no payments flow, no user
accounts, no trip history. The product surface is intentionally minimal so
that all the engineering effort goes into the parts that are actually hard:
spatial search, graph algorithms, concurrency-safe dispatch, and clean
system design. The React frontend exists only to visualize what the
backend computes.

> **Status:** Phases 1–2 complete (project scaffolding + foundational
> models, config, and the `/health` endpoint). Spatial search, routing,
> dispatch, concurrency, and benchmarking land in subsequent phases — see
> [Current Capabilities](#current-capabilities) and
> [Future Improvements](#future-improvements) below.

---

## Motivation

Ride-hailing dispatch is a genuinely interesting systems problem: it
combines nearest-neighbor search over a moving dataset, shortest-path
routing over a real road network, and concurrency control to avoid
double-booking a driver — all under a tight latency budget. Most take-home
and portfolio projects skip straight to CRUD and a database. This project
does the opposite: no database, no ORM, no web framework — just the
algorithms and the concurrency model that make dispatch work, implemented
and benchmarked from first principles.

---

## Features

- **Three interchangeable spatial search strategies** — Brute Force,
  Geohash, and KD-Tree — all implementing a single `spatial.Index`
  interface, so the dispatch engine never depends on a specific
  implementation.
- **Two routing algorithms** — Dijkstra and A* (with an admissible
  Haversine-distance heuristic) — over an in-memory adjacency-list graph.
- **A dispatch pipeline** that shortlists nearby drivers via spatial
  search, then runs real road-network routing to pick the driver with the
  minimum ETA (not just the closest one as the crow flies).
- **Concurrency-safe driver state**, using `sync.RWMutex` to allow
  concurrent reads (dispatch queries) while serializing writes (driver
  location updates), and to prevent double-assignment of a driver under
  concurrent dispatch requests.
- **A benchmark suite** comparing all spatial strategies and both routing
  algorithms on latency and search-quality metrics.
- **A lightweight React + Leaflet frontend** to visualize drivers, the
  rider, the selected driver, and the computed route on a live map.

---

## Architecture

```
Browser (React + Leaflet)
        |
        | HTTP / JSON
        v
Go backend (net/http, in-memory state, no database)
        |
        +-- internal/api          HTTP handlers
        +-- internal/dispatch     orchestrates spatial search + routing
        +-- internal/spatial      Index interface + bruteforce/geohash/kdtree
        +-- internal/routing      Dijkstra, A*
        +-- internal/graph        road-network graph construction
        +-- internal/models       shared types (Driver, Rider, Graph, Route)
        +-- internal/simulation   synthetic driver/rider generation
        +-- internal/benchmark    latency/quality comparison harness
        +-- internal/config       centralized tunables
        +-- internal/utils        Haversine distance, IDs, error types
```

The dispatch engine is written against the `spatial.Index` interface, not
a concrete implementation — Brute Force, Geohash, and KD-Tree are drop-in
replacements for each other. Full rationale in
[`docs/architecture.md`](docs/architecture.md) and
[`docs/design.md`](docs/design.md).

## Dispatch Pipeline

```
Rider Request
     │
     ▼
Spatial Index  (shortlist nearby drivers)
     │
     ▼
Top-K Candidate Drivers
     │
     ▼
A* Routing   (Haversine-distance heuristic)
     │
     ▼
Compare ETA → select minimum
     │
     ▼
Dispatch Response
```

---

## Algorithms Used

| Algorithm            | Purpose                                             | Status    |
|-----------------------|------------------------------------------------------|-----------|
| Haversine Distance     | Great-circle distance; spatial metric + A* heuristic | Implemented |
| Brute Force Search     | Baseline correctness reference for K-nearest search  | Planned (Phase 3) |
| Geohashing              | Bucketed candidate retrieval via neighbor expansion   | Planned (Phase 4) |
| KD-Tree                 | Pruned nearest/K-nearest neighbor search              | Planned (Phase 5) |
| Dijkstra                | Shortest path, uniform-cost expansion                 | Planned (Phase 7) |
| A*                      | Shortest path, heuristic-guided expansion             | Planned (Phase 8) |

## Data Structures

- **Adjacency-list graph** (`models.Graph`) for the road network — sparse
  by nature, so an adjacency list keeps memory and traversal cost
  proportional to edge count rather than `O(nodes²)`.
- **KD-Tree** (Phase 5) for spatially-partitioned nearest-neighbor search
  with branch pruning.
- **Geohash buckets** (Phase 4), a hash-map of geohash-string → drivers,
  for constant-time candidate bucket lookup plus neighbor-cell expansion.
- **Priority queue (binary heap)** (Phase 7) for Dijkstra/A* frontier
  management.

---

## Folder Structure

```
GeoDispatch-Engine/
├── backend/
│   ├── cmd/server/            entry point
│   ├── internal/
│   │   ├── api/                HTTP handlers, routing, middleware
│   │   ├── benchmark/          comparison harness (Phase 12)
│   │   ├── config/             centralized configuration
│   │   ├── dispatch/           dispatch orchestration (Phase 9)
│   │   ├── graph/              road-network graph loader (Phase 6)
│   │   ├── models/             shared data types
│   │   ├── routing/            Dijkstra + A* (Phases 7-8)
│   │   ├── simulation/         synthetic data generation (Phase 11)
│   │   ├── spatial/            Index interface + bruteforce/geohash/kdtree
│   │   └── utils/              Haversine, ID generation, error types
│   └── tests/                  integration tests
├── frontend/
│   └── src/
│       ├── components/         Navbar, MapView, SidePanel
│       ├── pages/               Dispatch Simulator, Benchmarks
│       ├── services/            backend API client
│       ├── hooks/                React hooks
│       └── types/                JSDoc type definitions
├── data/
│   ├── osm/                     preprocessed OpenStreetMap extracts
│   └── sample/                  small datasets for dev/tests
├── docs/
│   ├── architecture.md
│   ├── design.md
│   └── benchmarks.md
└── scripts/
```

---

## API Endpoints

| Method | Path                     | Description                                   | Status |
|--------|---------------------------|------------------------------------------------|--------|
| GET    | `/health`                  | Liveness/uptime check                          | Implemented |
| POST   | `/drivers`                  | Register a new driver                          | Planned |
| PUT    | `/drivers/{id}/location`     | Update a driver's location                     | Planned |
| DELETE | `/drivers/{id}`              | Remove a driver                                | Planned |
| POST   | `/dispatch`                   | Find the best driver for a rider location      | Planned |
| GET    | `/benchmark`                   | Retrieve spatial/routing benchmark results     | Planned |

---

## Complexity Analysis

| Operation                         | Brute Force | Geohash (typical)   | KD-Tree (balanced) |
|------------------------------------|-------------|-----------------------|----------------------|
| Insert driver                        | O(1)        | O(1)                   | O(log n)              |
| K-Nearest search                      | O(n log k)  | O(bucket size + neighbors) | O(√n + k) typical |
| Remove driver                          | O(n)        | O(1)                    | O(log n)              |

| Routing algorithm | Time complexity            | Notes |
|---------------------|------------------------------|-------|
| Dijkstra              | O((V + E) log V) with a binary heap | Explores uniformly outward from the source |
| A*                     | O((V + E) log V) worst case, typically far fewer nodes explored | Heuristic-guided toward the goal; optimal because Haversine distance is admissible |

Full reasoning behind each choice — including why Haversine distance is
guaranteed admissible for A* — lives in
[`docs/design.md`](docs/design.md).

---

## How To Run

### Prerequisites

- Go 1.22+
- Node.js 18+ and npm

### Backend

```bash
cd backend
go run ./cmd/server
# or, via the Makefile from the repo root:
make run-backend
```

The server starts on `:8080`. Verify it's up:

```bash
curl http://localhost:8080/health
```

Run the backend test suite:

```bash
cd backend
go test ./...
# or: make test-backend
```

### Frontend

```bash
cd frontend
npm install
npm run dev
# or, from the repo root: make install-frontend && make run-frontend
```

The dev server starts on `:5173` and proxies `/api/*` requests to the
backend on `:8080`.

---

## Current Capabilities

- Backend builds and serves `GET /health` over `net/http`.
- Core domain models (`Driver`, `Rider`, `Graph`, `Node`, `Edge`, `Route`,
  dispatch request/response types) are defined and unit-tested.
- `HaversineDistance` — the great-circle distance function used
  throughout spatial search and later as the A* heuristic — is
  implemented and unit-tested.
- Centralized configuration (`internal/config`) for all future tunables
  (candidate count, search radius, geohash precision, etc.).
- The `spatial.Index` interface is defined, establishing the contract
  Brute Force, Geohash, and KD-Tree will implement.
- Frontend scaffold: navbar, routing between the Dispatch Simulator and
  Benchmarks pages, a Leaflet map rendering base OpenStreetMap tiles, a
  side panel with placeholder dispatch stats, and an API service layer.

## Future Improvements

- Implement Brute Force, Geohash, and KD-Tree spatial search (Phases 3–5).
- Load a real road-network graph from preprocessed OpenStreetMap data
  (Phase 6).
- Implement Dijkstra and A* routing (Phases 7–8).
- Implement the dispatch engine tying spatial search and routing together
  (Phase 9).
- Wire the frontend up to real dispatch/simulation endpoints, including
  live markers and route polylines (Phase 10).
- Add concurrency safety for concurrent dispatch and driver updates, with
  double-assignment prevention (Phase 11).
- Build the benchmark suite and populate `docs/benchmarks.md` and the
  Benchmarks page with real numbers (Phase 12).
- Final documentation pass and polish, including real screenshots
  (Phase 13).

---

## Screenshots

_Placeholder — screenshots will be added once the Dispatch Simulator and
Benchmarks pages are wired to live backend data (Phase 10 onward)._

---

## License

MIT — see [`LICENSE`](LICENSE).
