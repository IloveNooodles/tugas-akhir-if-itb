import type { CreateRepositorySchema, Repository } from '~/types/repository';
import type { Response } from '~/types/response';

function transformGetRepositoryList(res: Response<Array<Repository>>) {
  return res.data;
}

export async function createRepository(
  body: CreateRepositorySchema,
  nuxtApp = useNuxtApp(),
) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/repositories`;

  return fetch(key, {
    method: 'POST',
    body: body,
    timeout: 1000,
  });
}

export async function getRepositoryList(nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/repositories`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetRepositoryList,
    server: false,
  });
}

export async function deleteRepositoryByID(id: string, nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/repositories/${id}`;

  return fetch(key, {
    method: 'DELETE',
    timeout: 500,
  });
}
