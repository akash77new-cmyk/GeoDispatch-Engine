// SidePanel is the control + readout surface next to the map on the
// Dispatch Simulator page. It currently renders placeholder stats and
// action buttons; it will be wired to real dispatch state in Phase 10.
export default function SidePanel() {
  return (
    <aside className="side-panel">
      <section>
        <h2>Dispatch Result</h2>
        <div className="stat-row">
          <span className="label">Selected Driver</span>
          <span className="value">—</span>
        </div>
        <div className="stat-row">
          <span className="label">ETA</span>
          <span className="value">—</span>
        </div>
        <div className="stat-row">
          <span className="label">Distance</span>
          <span className="value">—</span>
        </div>
        <div className="stat-row">
          <span className="label">Algorithm Used</span>
          <span className="value">—</span>
        </div>
        <div className="stat-row">
          <span className="label">Dispatch Time</span>
          <span className="value">—</span>
        </div>
      </section>

      <section>
        <h2>Candidate Drivers</h2>
        <p className="empty-note">
          No candidates yet. Generate drivers and a rider, then dispatch.
        </p>
      </section>

      <section className="button-row">
        <h2>Actions</h2>
        <button className="btn" type="button" disabled>
          Generate Drivers
        </button>
        <button className="btn" type="button" disabled>
          Generate Rider
        </button>
        <button className="btn btn-primary" type="button" disabled>
          Dispatch
        </button>
        <p className="empty-note">
          Actions are wired up in Phase 10 once the dispatch engine and
          simulation endpoints exist on the backend.
        </p>
      </section>
    </aside>
  )
}
