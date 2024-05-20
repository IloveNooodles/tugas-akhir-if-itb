import type { DeployResponse } from '~/types/deployment';
import type { DoRemoteDeploySchema } from '~/types/remote';

export async function Deploy(
  body: DoRemoteDeploySchema,
  nuxtApp = useNuxtApp(),
) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/deployments/deploy`;

  const response = await fetch<DeployResponse>(key, {
    method: 'POST',
    body: body,
  });

  return response;
}

export async function DeleteDeploy(ids: string[], nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/deployments/deploy/delete`;

  const response = await fetch(key, {
    method: 'POST',
    body: {
      deployment_ids: ids,
    },
  });

  return response;
}
