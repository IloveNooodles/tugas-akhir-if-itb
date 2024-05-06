import type { UserLoginResponse, UserLoginSchema } from '~/types/user';

export async function login(body: UserLoginSchema, nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const response = await fetch<UserLoginResponse>('/api/v1/users/login', {
    method: 'POST',
    body: body,
    timeout: 500,
  });

  return response;
}

export async function refresh(nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const response = await fetch<UserLoginResponse>('/api/v1/users/refresh', {
    method: 'POST',
    timeout: 500,
  });

  return response;
}
