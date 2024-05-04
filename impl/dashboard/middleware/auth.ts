const isAuthenticated = () => {
  const accessToken = useCookie('accessToken');
  const refreshToken = useCookie('refreshToken');

  // console.log('AT, RT', accessToken.value, refreshToken.value);

  if (!accessToken.value || !refreshToken.value) return false;

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

  // if (!isAuthenticated()) {
  //   return navigateTo('/login');
  // }
});
