import { Routes, Route } from 'react-router-dom'
import Navbar from './components/Navbar.jsx'
import DispatchSimulatorPage from './pages/DispatchSimulatorPage.jsx'
import BenchmarksPage from './pages/BenchmarksPage.jsx'

// App is the top-level layout: a persistent navbar plus the two pages
// described in the project brief. Routing is intentionally flat (no
// nested routes) since the frontend is a thin visualization layer, not
// the product itself.
export default function App() {
  return (
    <div className="app-shell">
      <Navbar />
      <main className="app-content">
        <Routes>
          <Route path="/" element={<DispatchSimulatorPage />} />
          <Route path="/benchmarks" element={<BenchmarksPage />} />
        </Routes>
      </main>
    </div>
  )
}
