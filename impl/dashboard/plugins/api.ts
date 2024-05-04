import { refresh } from '~/api/auth';

export default defineNuxtPlugin(async (nuxtApp) => {
  const config = useRuntimeConfig();

  async function createApiClient() {
    const {
      public: { baseURL, apiKey, adminApiKey },
    } = config;

    const instance = $fetch.create({
      baseURL,
      headers: {
        ['X-API-Key']: apiKey,
        ['X-Admin-API-Key']: adminApiKey ?? '', // Optional admin key
      },
      async onResponseError({ options, request, response }) {
        if (response.status === 401) {
          try {
            await refresh();
            // Update headers or store tokens as needed
          } catch (err) {
            console.error('API refresh error:', err);
            // Handle refresh error gracefully (e.g., logout)
          }
        }
      },
      credentials: 'include',
      retry: 1,
      retryDelay: 1000,
      retryStatusCodes: [401, 408, 409, 425, 429, 500, 502, 503, 504],
    });

    return instance;
  }

  const apiInstance = await createApiClient();

  return {
    provide: {
      api: apiInstance,
    },
  };
});
