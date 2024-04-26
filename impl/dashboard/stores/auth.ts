import { defineStore } from 'pinia';

export const useAuthStore = defineStore('counter', () => {
  const token = ref('');
  function get() {
    return token.value;
  }

  function set(t: string) {
    token.value = t;
  }

  return { token, get, set };
});
