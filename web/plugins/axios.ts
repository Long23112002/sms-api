import axios from 'axios'

// Function to get API base URL from runtime config
function getApiBaseURL(): string {
  // Try to get from Nuxt runtime config (available in client after app loads)
  if (typeof window !== 'undefined') {
    // Check window.__NUXT__ which contains the runtime config
    const nuxtConfig = (window as any).__NUXT__?.config?.publicRuntimeConfig
    if (nuxtConfig?.apiBaseURL) {
      return nuxtConfig.apiBaseURL
    }
    // Also try $nuxt instance if available
    if ((window as any).$nuxt?.$config?.publicRuntimeConfig?.apiBaseURL) {
      return (window as any).$nuxt.$config.publicRuntimeConfig.apiBaseURL
    }
  }
  // Fallback to environment variable (for build time or SSR)
  return process.env.API_BASE_URL || 'http://103.77.160.47:8000'
}

const client = axios.create({
  baseURL: getApiBaseURL(),
  headers: {
    'X-Client-Version': process.env.GITHUB_SHA || 'dev',
  },
})

// Update baseURL when runtime config is available (for client-side)
if (typeof window !== 'undefined') {
  // Function to update baseURL from runtime config
  const updateBaseURL = () => {
    const newBaseURL = getApiBaseURL()
    if (newBaseURL && newBaseURL !== client.defaults.baseURL) {
      client.defaults.baseURL = newBaseURL
    }
  }

  // Try to update immediately
  updateBaseURL()

  // Also update when DOM is ready (Nuxt config should be available by then)
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', updateBaseURL)
  } else {
    // DOM already loaded, try after a short delay to ensure Nuxt is ready
    setTimeout(updateBaseURL, 100)
  }

  // Also listen for Nuxt ready event if available
  if ((window as any).$nuxt) {
    ;(window as any).$nuxt.$once('hook:mounted', updateBaseURL)
  }
}

// Storage keys
const API_KEY_STORAGE_KEY = 'api_key'
const AUTH_USER_STORAGE_KEY = 'authUser'

// Restore API key from localStorage on initialization
if (process.client) {
  const savedApiKey = localStorage.getItem(API_KEY_STORAGE_KEY)
  if (savedApiKey) {
    client.defaults.headers.common['x-api-key'] = savedApiKey
  }
}

export function setAuthHeader(token: string | null) {
  if (token) {
  client.defaults.headers.common.Authorization = 'Bearer ' + token
  } else {
    delete client.defaults.headers.common.Authorization
  }
}

export function setApiKey(apiKey: string | null) {
  if (process.client) {
    if (apiKey) {
      localStorage.setItem(API_KEY_STORAGE_KEY, apiKey)
      client.defaults.headers.common['x-api-key'] = apiKey
    } else {
      localStorage.removeItem(API_KEY_STORAGE_KEY)
      delete client.defaults.headers.common['x-api-key']
    }
  } else {
  client.defaults.headers.common['x-api-key'] = apiKey ?? ''
  }
}

export function getApiKey(): string | null {
  if (process.client) {
    return localStorage.getItem(API_KEY_STORAGE_KEY)
  }
  return null
}

export function saveAuthUser(authUser: any) {
  if (process.client) {
    if (authUser) {
      localStorage.setItem(AUTH_USER_STORAGE_KEY, JSON.stringify(authUser))
    } else {
      localStorage.removeItem(AUTH_USER_STORAGE_KEY)
    }
  }
}

export function getAuthUser(): any | null {
  if (process.client) {
    const saved = localStorage.getItem(AUTH_USER_STORAGE_KEY)
    if (saved) {
      try {
        return JSON.parse(saved)
      } catch (e) {
        return null
      }
    }
  }
  return null
}

export default client
