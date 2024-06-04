import type { History } from '~/types/history';
import type { Response } from '~/types/response';

function transformGetDeploymentHistoryList(res: Response<Array<History>>) {
  return res.data;
}

export async function getDeploymentHistoryList(nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/histories`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetDeploymentHistoryList,
    server: false,
  });
}

export async function getDeploymentHistoryListWithDeployments(
  nuxtApp = useNuxtApp(),
  id: string,
) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/histories?deployment_id=${id}`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetDeploymentHistoryList,
    server: false,
  });
}

export async function getDeploymentHistoryListWithGroups(
  nuxtApp = useNuxtApp(),
  id: string,
) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/histories?group_id=${id}`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetDeploymentHistoryList,
    server: false,
  });
}

export async function getDeploymentHistoryListWithDevices(
  nuxtApp = useNuxtApp(),
  id: string,
) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/histories?device_id=${id}`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetDeploymentHistoryList,
    server: false,
  });
}

export async function getDeploymentHistoryListWithRepository(
  nuxtApp = useNuxtApp(),
  id: string,
) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/histories?repository_id=${id}`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetDeploymentHistoryList,
    server: false,
  });
}
