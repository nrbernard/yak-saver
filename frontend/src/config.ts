// API configuration
// In development: uses localhost:8080
// In production: uses empty string for same-origin requests
export const API_BASE_URL = import.meta.env.VITE_API_BASE_URL !== undefined
  ? import.meta.env.VITE_API_BASE_URL
  : 'http://localhost:8080';
