import { useAuthStore } from '#imports';

export const useAPI = () => {
  const config = useRuntimeConfig();
  const instance = $fetch.create({
    baseURL: config.public.baseURL,
    headers: {
      Authorization: `Bearer ${useAuthStore().get()}`,
      ['X-API-Key']: config.public.apiKey,
      ['X-Admin-API-Key']: config.public.adminApiKey,
    },
  });

  return instance;
};
