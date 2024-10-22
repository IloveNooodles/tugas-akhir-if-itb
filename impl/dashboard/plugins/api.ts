import { callWithNuxt } from '#app';
import { refresh } from '~/api/auth';

export default defineNuxtPlugin(async (nuxtApp) => {
  const config = useRuntimeConfig();
  const nxtApp = useNuxtApp();

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
      async onResponseError({ error, request, response }) {
        if (response.status === 401) {
          try {
            const response = await refresh(nxtApp);
            useLocalStorage('accessToken', '').value =
              response.data.access_token;
            useLocalStorage('refreshToken', '').value =
              response.data.refresh_token;
          } catch (err) {
            useLocalStorage('accessToken', '').value = '';
            useLocalStorage('refreshToken', '').value = '';

            callWithNuxt(nxtApp, () => navigateTo('/login'));
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
