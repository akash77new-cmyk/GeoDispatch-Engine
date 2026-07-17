// BenchmarksPage is Page 2 from the project brief: comparison tables for
// spatial search strategies and routing strategies. Populated with real
// numbers from the backend's /benchmark endpoint in Phase 12; for now it
// renders the table structure with placeholder rows.
const spatialRows = [
  { name: 'Brute Force', avgMs: '—', medianMs: '—', p95Ms: '—' },
  { name: 'Geohash', avgMs: '—', medianMs: '—', p95Ms: '—' },
  { name: 'KD-Tree', avgMs: '—', medianMs: '—', p95Ms: '—' },
]

const routingRows = [
  { name: 'Dijkstra', execMs: '—', nodesExplored: '—', pathLength: '—' },
  { name: 'A*', execMs: '—', nodesExplored: '—', pathLength: '—' },
]

export default function BenchmarksPage() {
  return (
    <div className="benchmarks-page">
      <h1>Benchmarks</h1>
      <p className="empty-note">
        Data will populate once the spatial search and routing engines
        (Phases 3–8) and the benchmark suite (Phase 12) are implemented.
      </p>

      <h2>Spatial Search</h2>
      <table className="benchmark-table">
        <thead>
          <tr>
            <th>Strategy</th>
            <th>Avg Query Time</th>
            <th>Median Query Time</th>
            <th>P95 Latency</th>
          </tr>
        </thead>
        <tbody>
          {spatialRows.map((row) => (
            <tr key={row.name}>
              <td>{row.name}</td>
              <td className="numeric">{row.avgMs}</td>
              <td className="numeric">{row.medianMs}</td>
              <td className="numeric">{row.p95Ms}</td>
            </tr>
          ))}
        </tbody>
      </table>

      <h2>Routing</h2>
      <table className="benchmark-table">
        <thead>
          <tr>
            <th>Algorithm</th>
            <th>Execution Time</th>
            <th>Nodes Explored</th>
            <th>Path Length</th>
          </tr>
        </thead>
        <tbody>
          {routingRows.map((row) => (
            <tr key={row.name}>
              <td>{row.name}</td>
              <td className="numeric">{row.execMs}</td>
              <td className="numeric">{row.nodesExplored}</td>
              <td className="numeric">{row.pathLength}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
