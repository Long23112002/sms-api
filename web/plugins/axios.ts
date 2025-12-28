import axios from 'axios'

const client = axios.create({
  baseURL: process.env.API_BASE_URL || 'http://localhost:8000',
  headers: {
    'X-Client-Version': process.env.GITHUB_SHA || 'dev',
  },
})

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
