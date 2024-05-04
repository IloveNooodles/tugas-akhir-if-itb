import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', () => {
  const auth = ref({
    accessToken: '',
    refreshToken: '',
  });

  function getAccessToken() {
    return auth.value.accessToken;
  }

  function getRefreshToken() {
    return auth.value.refreshToken;
  }

  function set(t: string, rt: string) {
    auth.value.accessToken = t;
    auth.value.refreshToken = rt;
  }

  return { auth, getAccessToken, getRefreshToken, set };
});
