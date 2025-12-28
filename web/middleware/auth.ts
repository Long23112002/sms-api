import { Context, Middleware } from '@nuxt/types'

const authMiddleware: Middleware = async (context: Context) => {
  // Wait for auth state to be restored from localStorage
  // This ensures the plugin has time to restore the session
  if (!context.store.getters.authStateChanged) {
    // Wait a bit for the plugin to restore state
    await new Promise((resolve) => setTimeout(resolve, 100))
    
    // If still not changed, wait a bit more (max 1 second)
    let attempts = 0
    while (!context.store.getters.authStateChanged && attempts < 10) {
      await new Promise((resolve) => setTimeout(resolve, 100))
      attempts++
    }
  }

  // Now check if user is authenticated
  if (context.store.getters.getAuthUser === null) {
    context.redirect('/login', { to: context.route.path })
  }
}

export default authMiddleware
