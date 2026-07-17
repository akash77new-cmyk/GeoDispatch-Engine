// api.js is the single place the frontend talks to the backend from.
// Centralizing fetch calls here (rather than scattering fetch() through
// components) means the base URL, error handling, and JSON parsing only
// need to be written once.

const BASE_URL = '/api'

async function request(path, options = {}) {
  const res = await fetch(`${BASE_URL}${path}`, {
    headers: { 'Content-Type': 'application/json' },
    ...options,
  })

  if (!res.ok) {
    const body = await res.json().catch(() => ({}))
    throw new Error(body.error || `Request failed: ${res.status}`)
  }

  return res.json()
}

// getHealth checks GET /health. Used by the app shell to verify backend
// connectivity; wired up further in later phases.
export function getHealth() {
  return request('/health')
}

// The following API calls are part of the target surface but are
// implemented once their backend endpoints exist:
//   createDriver(driver)         -> POST /drivers          (Phase 9/10)
//   updateDriverLocation(id, loc)-> PUT  /drivers/{id}/location
//   removeDriver(id)             -> DELETE /drivers/{id}
//   dispatch(request)            -> POST /dispatch
//   getBenchmarks()              -> GET  /benchmark
