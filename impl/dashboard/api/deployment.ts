import type { CreateDeploymentSchema, Deployment } from '~/types/deployment';
import type { Response } from '~/types/response';

function transformGetDeploymentList(res: Response<Deployment[]>) {
  return res.data;
}

function transformGetDeploymentDetailByID(res: Response<Deployment>) {
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

export async function getDeploymentDetailByID(
  id: string,
  nuxtApp = useNuxtApp(),
) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/deployments/${id}`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetDeploymentDetailByID,
    server: false,
  });
}

export async function deleteDeploymentByID(id: string, nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/deployments/${id}`;

  return fetch(key, {
    method: 'DELETE',
    timeout: 500,
  });
}

export async function createDeployment(
  body: CreateDeploymentSchema,
  nuxtApp = useNuxtApp(),
) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/deployments`;

  return fetch(key, {
    method: 'POST',
    timeout: 1000,
    body: body,
  });
}
