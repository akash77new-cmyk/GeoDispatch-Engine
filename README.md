# 🌍 GeoRoute Engine: Geospatial Driver Discovery & ETA Optimization

GeoRoute Engine is a backend systems project inspired by the core challenge faced by modern ride-hailing platforms: **efficiently matching riders with the most suitable drivers**.

Instead of selecting the geographically closest driver using straight-line distance, GeoRoute Engine combines **spatial indexing** with **graph-based routing** to dispatch the driver with the lowest **Estimated Time of Arrival (ETA)**. The project showcases how computational geometry, graph algorithms, concurrent programming, and geospatial data can be combined to build an intelligent dispatch system.

---

# 🎯 Project Motivation

Real-world ride-hailing applications cannot rely solely on Euclidean distance to assign drivers. Two drivers may be equally close to a rider, but road layouts, one-way streets, and traffic conditions can result in significantly different arrival times.

GeoRoute Engine addresses this by implementing a **two-stage dispatch pipeline**:

1. Efficiently shortlist nearby drivers using spatial indexing.
2. Compute the optimal route over a road network and dispatch the driver with the minimum ETA.

This project focuses on the algorithmic and backend engineering aspects of geospatial routing systems.

---

# ✨ Features

- 🚖 Two-stage driver dispatch pipeline
- 📍 Multiple spatial indexing techniques
  - Brute Force
  - Geohashing
  - KD-Tree
- 🗺️ Road-network routing over OpenStreetMap (OSM) data
- 🧠 Manual implementation of Dijkstra's Algorithm and A* Search
- ⚡ Concurrent Go REST APIs using goroutines and mutexes
- 🌐 Interactive React + Leaflet visualization
- 📊 Built-in benchmarking framework for algorithm comparison

---

# 🏗️ Dispatch Pipeline

```text
Driver Locations
        │
        ▼
Spatial Index
(Brute Force / Geohash / KD-Tree)
        │
        ▼
Top-K Candidate Drivers
        │
        ▼
Road Network Routing
(Dijkstra / A*)
        │
        ▼
ETA Calculation
        │
        ▼
Optimal Driver Selection
        │
        ▼
Dispatch Response
```

---

# 🏛️ System Architecture

```text
                  React + Leaflet
                        │
                 REST API Requests
                        │
                        ▼
               ┌────────────────────┐
               │     Go Backend     │
               └────────────────────┘
                  │       │        │
                  ▼       ▼        ▼
           Spatial     Routing   Benchmarking
            Search      Engine
      (KD-Tree /      (Dijkstra /
       Geohash)          A*)
                  │
                  ▼
           Driver Dispatch
```

---

# 🧩 Algorithms Implemented

## 📍 Spatial Search

Implemented from scratch:

- Brute Force Search
- Geohashing
- KD-Tree

Supports:

- Driver insertion
- Driver updates
- Driver removal
- K-nearest driver search

---

## 🗺️ Routing Engine

Implemented from scratch:

- Dijkstra's Algorithm
- A* Search (using the Haversine-distance heuristic)

Supports:

- Shortest-path computation
- Route generation
- ETA estimation

---

## ⚡ Concurrent Backend

- Goroutines
- Mutex-based synchronization
- Thread-safe driver management
- Duplicate dispatch prevention

---

# 📊 Benchmarking

GeoRoute Engine compares multiple approaches to evaluate performance under different workloads.

### Spatial Search Comparison

Algorithms:

- Brute Force
- Geohash
- KD-Tree

Metrics:

- Search latency
- Driver lookup time
- Memory usage
- Scalability

### Routing Comparison

Algorithms:

- Dijkstra
- A*

Metrics:

- Execution time
- Nodes explored
- Route length
- ETA computation time

---

# 💻 Tech Stack

### Backend

- Go
- REST API
- Goroutines
- Mutexes

### Frontend

- React
- Vite
- Leaflet

### Algorithms

- KD-Tree
- Geohashing
- Dijkstra's Algorithm
- A* Search
- Haversine Distance

### Data

- OpenStreetMap (OSM)

---

# 📂 Project Structure

```text
GeoRoute-Engine/

backend/
├── cmd/
├── internal/
│   ├── api/
│   ├── benchmark/
│   ├── config/
│   ├── dispatch/
│   ├── graph/
│   ├── models/
│   ├── routing/
│   ├── simulation/
│   ├── spatial/
│   │   ├── bruteforce/
│   │   ├── geohash/
│   │   └── kdtree/
│   └── utils/
├── tests/
└── go.mod

frontend/
├── public/
└── src/
    ├── components/
    ├── hooks/
    ├── pages/
    ├── services/
    └── types/

data/
docs/
scripts/

README.md
```

---

# 🖥️ Visualization

The React + Leaflet interface allows users to:

- 📍 Generate driver locations
- 🚶 Generate rider requests
- 🚗 Display candidate drivers
- 🛣️ Visualize computed routes
- 📈 Compare routing algorithms
- 📊 Benchmark spatial indexing techniques
- ⚡ Observe dispatch decisions in real time

---

# 🚀 Future Improvements

Potential enhancements include:

- 🚦 Traffic-aware ETA estimation using live traffic conditions.
- 🌍 Support for multiple cities through additional OpenStreetMap datasets.
- 🚗 Dynamic driver simulation with continuous location updates.
- 📍 Advanced spatial indexing using structures such as R-Trees or Quad Trees.
- ⚡ Route caching to reduce repeated shortest-path computations.
- 📊 Enhanced benchmarking with throughput, latency, memory usage, and scalability analysis.
- 🔄 Live dispatch simulation using WebSockets.
- 🗄️ Persistent storage using PostgreSQL.
- 🐳 Containerized deployment using Docker.
- ☁️ Migration to a microservice-based architecture by separating routing, dispatch, and spatial indexing into independent services.

---

# 📚 Learning Outcomes

This project demonstrates practical implementation of:

- Backend Systems Engineering
- Geospatial Computing
- Graph Algorithms
- Computational Geometry
- Concurrent Programming in Go
- REST API Design
- Performance Benchmarking
- Interactive Data Visualization

---

## ⭐ If you found this project interesting, consider giving it a star!
