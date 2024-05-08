import type { Response } from '~/types/response';
import type { User } from '~/types/user';

function transformGetUserDetail(res: Response<User>) {
  return res.data;
}

function transformGetUserLists(res: Response<Array<User>>) {
  return res.data;
}

export async function getUserDetail(id: string, nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/users/${id}`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetUserDetail,
    server: false,
  });
}

export async function getUserLists(nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;

  return useLazyFetch('/api/v1/users', {
    $fetch: fetch,
    transform: transformGetUserLists,
    server: false,
  });
}

export async function deleteUserByID(id: string, nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/users/${id}`;

  return fetch(key, {
    method: 'DELETE',
    timeout: 500,
  });
}
