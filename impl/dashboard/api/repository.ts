import type { Repository } from '~/types/repository';
import type { Response } from '~/types/response';

function transformGetRepositories(res: Response<Array<Repository>>) {
  return res.data;
}

export async function getRepositories(id: string, nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/users/${id}`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetRepositories,
    server: false,
  });
}
