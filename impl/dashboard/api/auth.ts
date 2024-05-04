import type { UserLoginResponse, UserLoginSchema } from '~/types/user';

export async function login(body: UserLoginSchema) {
  const nuxtApp = useNuxtApp();
  const fetch = nuxtApp.$api;
  const response = await fetch<UserLoginResponse>('/api/v1/users/login', {
    method: 'POST',
    body: body,
  });

  return response;
}

export async function refresh() {
  const nuxtApp = useNuxtApp();
  const fetch = nuxtApp.$api;
  const response = await fetch<UserLoginResponse>('/api/v1/users/refresh', {
    method: 'POST',
  });

  return response;
}
