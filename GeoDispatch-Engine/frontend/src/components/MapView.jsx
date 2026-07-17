import { MapContainer, TileLayer } from 'react-leaflet'

// MapView renders the Leaflet map that the Dispatch Simulator page sits
// on. In Phase 2 it only renders base OpenStreetMap tiles; driver/rider
// markers, the selected-driver highlight, and the computed route
// polyline are added in Phase 10 once the backend can produce them.
//
// Default center is New Delhi, matching the sample/simulated data this
// project will generate.
const DEFAULT_CENTER = [28.6139, 77.209]
const DEFAULT_ZOOM = 13

export default function MapView() {
  return (
    <MapContainer
      center={DEFAULT_CENTER}
      zoom={DEFAULT_ZOOM}
      style={{ height: '100%', width: '100%' }}
    >
      <TileLayer
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
    </MapContainer>
  )
}
