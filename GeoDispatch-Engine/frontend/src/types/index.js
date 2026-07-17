// This project uses plain JS (not TypeScript) to keep the frontend
// lightweight, per the project brief. These JSDoc typedefs document the
// shapes exchanged with the backend so editors still get autocomplete
// and basic type-checking without a build-time TS toolchain.

/**
 * @typedef {Object} Location
 * @property {number} lat
 * @property {number} lng
 */

/**
 * @typedef {Object} Driver
 * @property {string} id
 * @property {Location} location
 * @property {'available'|'assigned'|'offline'} status
 * @property {string} updated_at
 */

/**
 * @typedef {Object} DispatchResult
 * @property {Driver} selected_driver
 * @property {number} eta_seconds
 * @property {number} distance_meters
 * @property {string[]} [route]
 * @property {string} algorithm_used
 * @property {Driver[]} candidate_drivers
 * @property {number} dispatch_time_ms
 */

export {}
