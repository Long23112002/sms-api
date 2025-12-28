import { Context } from '@nuxt/types'
import { getApiKey, getAuthUser, setApiKey } from '~/plugins/axios'

export default function ({ store }: Context) {
  // Only run on client side
  if (process.client) {
    // Restore API key from localStorage
    const savedApiKey = getApiKey()
    if (savedApiKey) {
      setApiKey(savedApiKey)
    }

    // Restore auth user from localStorage
    const savedAuthUser = getAuthUser()
    if (savedAuthUser) {
      // Set auth state changed to true so pages know auth is ready
      store.commit('setAuthUser', savedAuthUser)
    } else {
      // If no saved auth user, mark auth state as changed anyway
      store.commit('setAuthStateChanged', true)
    }
  }
}

