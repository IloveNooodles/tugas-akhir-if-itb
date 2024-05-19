// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: {
    enabled: true,
    timeline: {
      enabled: true,
    },
  },
  modules: ['@nuxt/ui', '@pinia/nuxt', '@vueuse/nuxt'],
  runtimeConfig: {
    public: {
      baseURL: 'http://localhost:8000/api/v1',
      apiKey: 'inicontohapikey',
      adminApiKey: 'inicontohadminapikey',
    },
  },
  app: {
    head: {
      title: 'Remote Deployment',
      charset: 'utf-8',
      viewport: 'width=device-width, initial-scale=1',
      meta: [
        {
          name: 'description',
          content: 'Remote deployment description',
        },
      ],
    },
    pageTransition: { name: 'page', mode: 'out-in' },
  },
  typescript: {
    typeCheck: false,
    strict: false,
  },
  css: ['~/assets/css/tailwind.css', '~/assets/css/index.css'],
});
