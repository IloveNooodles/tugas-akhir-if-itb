// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxt/ui', '@pinia/nuxt'],
  runtimeConfig: {
    public: {
      baseURL: 'http://localhost:8000/api/v1',
      apiKey: 'inicontohapikey',
      adminApiKey: 'inicontohadminapikey',
    },
  },
});
