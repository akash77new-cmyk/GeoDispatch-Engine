import MapView from '../components/MapView.jsx'
import SidePanel from '../components/SidePanel.jsx'

// DispatchSimulatorPage is Page 1 from the project brief: a live map
// where drivers/riders are generated and dispatch requests are sent to
// the backend for visualization. Interactivity is added in Phase 10.
export default function DispatchSimulatorPage() {
  return (
    <div className="dispatch-page">
      <div className="map-pane">
        <MapView />
      </div>
      <SidePanel />
    </div>
  )
}
