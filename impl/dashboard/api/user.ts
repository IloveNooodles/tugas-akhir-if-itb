import type { Response } from '~/types/response';
import type { User } from '~/types/user';

export async function getUserDetail(id: string) {
  const nuxtApp = useNuxtApp();
  const fetch = nuxtApp.$api;
  const key = `/api/v1/users/${id}`;

  return useLazyFetch<Response<User>>(key, {
    $fetch: fetch,
  });
}

export async function getUserLists() {
  const nuxtApp = useNuxtApp();
  const fetch = nuxtApp.$api;

  return useLazyFetch<Response<Array<User>>>('/api/v1/users', {
    $fetch: fetch,
  });
}
