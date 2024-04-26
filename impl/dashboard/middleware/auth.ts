const isAuthenticated = () => {
  const token = useAuthStore().get();
  const cfg = useRuntimeConfig();
  const apiToken = cfg.public.apiToken;

  if (!token) return false;
  if (token != apiToken) return false;

  return true;
};

export default defineNuxtRouteMiddleware((to, from) => {
  const nuxtApp = useNuxtApp();

  if (
    import.meta.client &&
    nuxtApp.isHydrating &&
    nuxtApp.payload.serverRendered
  )
    return;
  if (!isAuthenticated()) {
    return navigateTo('/login');
  }
});
