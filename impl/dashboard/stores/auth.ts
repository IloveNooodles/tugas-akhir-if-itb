import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', () => {
  const accessToken = ref('')
  const refreshToken = ref('')

  return { accessToken, refreshToken };
});
