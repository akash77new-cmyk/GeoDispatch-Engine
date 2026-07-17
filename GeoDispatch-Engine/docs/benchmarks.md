# Benchmarks

No benchmarks have been run yet — this document is a placeholder that
will be populated in **Phase 12 (Benchmark Suite)**, after the spatial
search strategies (Phases 3–5) and routing algorithms (Phases 7–8) exist.

## Planned Comparisons

### Spatial Search

| Strategy     | Avg Query Time | Median Query Time | P95 Latency |
|--------------|----------------|--------------------|-------------|
| Brute Force  | TBD            | TBD                | TBD         |
| Geohash      | TBD            | TBD                | TBD         |
| KD-Tree      | TBD            | TBD                | TBD         |

### Routing

| Algorithm | Execution Time | Nodes Explored | Path Length |
|-----------|-----------------|-----------------|-------------|
| Dijkstra  | TBD             | TBD             | TBD         |
| A*        | TBD             | TBD             | TBD         |

## Methodology (planned)

- Spatial search strategies will be benchmarked against the same
  driver dataset and the same set of query points, so results are
  directly comparable.
- Routing algorithms will be benchmarked against the same graph and the
  same set of (origin, destination) pairs.
- The `GET /benchmark` endpoint will expose these numbers so the
  frontend's Benchmarks page can render them without a separate offline
  step.
