import { useAuthStore } from '#imports';

export const useAPI = () => {
  const config = useRuntimeConfig();
  const { get, set } = useAuthStore();
  const instance = $fetch.create({
    baseURL: config.public.baseURL,
    headers: {
      Authorization: `Bearer ${get()}`,
      ['X-API-Key']: config.public.apiKey,
      ['X-Admin-API-Key']: config.public.adminApiKey,
    },
    onResponseError({ options, request, response }) {
      if (response.status === 401) {
        set('');
      }
    },
  });

  return instance;
};
