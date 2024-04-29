export default defineNuxtRouteMiddleware((to, from) => {
  const nuxtApp = useNuxtApp();
  const token = useAuthStore().get();

  if (
    import.meta.client &&
    nuxtApp.isHydrating &&
    nuxtApp.payload.serverRendered
  )
    return;
  if (!token) {
    return navigateTo('/login');
  }
});
