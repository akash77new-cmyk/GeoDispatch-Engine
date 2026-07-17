import { useEffect, useState } from 'react'
import { getHealth } from '../services/api.js'

// useBackendHealth polls the backend's /health endpoint once on mount
// and reports whether it is reachable. It exists as a small proof that
// the frontend/backend wiring works end-to-end after Phase 2, ahead of
// the real dispatch/simulation hooks added in Phase 10.
export function useBackendHealth() {
  const [status, setStatus] = useState('checking')

  useEffect(() => {
    let cancelled = false

    getHealth()
      .then(() => {
        if (!cancelled) setStatus('online')
      })
      .catch(() => {
        if (!cancelled) setStatus('offline')
      })

    return () => {
      cancelled = true
    }
  }, [])

  return status
}
