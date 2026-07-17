import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// GeoDispatch Engine frontend build config.
// The dev server proxies /api requests to the Go backend so the browser
// can call relative paths without hardcoding a host/port or dealing with
// CORS during local development.
export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
})
