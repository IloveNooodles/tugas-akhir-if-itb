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
