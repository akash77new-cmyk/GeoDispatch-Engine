import { NavLink } from 'react-router-dom'

// Navbar is the persistent top-level navigation between the two pages
// this frontend has: the live dispatch simulator and the algorithm
// benchmarks view.
export default function Navbar() {
  return (
    <header className="navbar">
      <div className="navbar-brand">
        <span className="mark">●</span>
        <span className="name">GeoDispatch Engine</span>
      </div>
      <nav className="navbar-links">
        <NavLink
          to="/"
          end
          className={({ isActive }) => 'navbar-link' + (isActive ? ' active' : '')}
        >
          Dispatch Simulator
        </NavLink>
        <NavLink
          to="/benchmarks"
          className={({ isActive }) => 'navbar-link' + (isActive ? ' active' : '')}
        >
          Benchmarks
        </NavLink>
      </nav>
    </header>
  )
}
