import type { History } from '~/types/history';
import type { Response } from '~/types/response';

function transformGetDeploymentList(res: Response<Array<History>>) {
  return res.data;
}

export async function GetDeploymentList(nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/histories`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetDeploymentList,
    server: false,
  });
}
