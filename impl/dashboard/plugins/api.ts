import { callWithNuxt } from '#app';
import { refresh } from '~/api/auth';

export default defineNuxtPlugin(async (nuxtApp) => {
  const config = useRuntimeConfig();
  const nxtApp = useNuxtApp()

  async function createApiClient() {
    const {
      public: { baseURL, apiKey, adminApiKey },
    } = config;

    const instance = $fetch.create({
      baseURL,
      headers: {
        ['X-API-Key']: apiKey,
        ['X-Admin-API-Key']: adminApiKey ?? '',
      },
      onResponseError({ error, request, response }) {
        if (response.status === 401) {
          try {
            refresh(nxtApp);
          } catch (err) {
            callWithNuxt(nxtApp, () =>
              navigateTo('/login', { redirectCode: 301 }),
            );
          }
        }
      },
      credentials: 'include',
      retry: 2,
      retryDelay: 2000,
      timeout: 500,
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
