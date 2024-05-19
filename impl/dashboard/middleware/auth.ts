import { useStorage } from '@vueuse/core';

const isAuthenticated = () => {
  const at = useStorage('accessToken', '');
  const rt = useStorage('refreshToken', '');

  if (!at.value || !rt.value) return false;

  return true;
};

export default defineNuxtRouteMiddleware((to, from) => {
  if (!isAuthenticated() && to.path !== '/login') {
    return navigateTo('/login');
  }
});
