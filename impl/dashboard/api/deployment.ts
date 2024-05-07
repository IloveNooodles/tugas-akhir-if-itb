import type { Deployment } from '~/types/deployment';
import type { Response } from '~/types/response';

function transformGetDeploymentList(res: Response<Deployment[]>) {
  return res.data;
}

export async function getDeploymentList(nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/deployments`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetDeploymentList,
    server: false,
  });
}
