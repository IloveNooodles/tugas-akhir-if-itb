export async function Deploy(ids: string[], nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/deployments/deploy`;

  return fetch(key, {
    method: 'POST',
    body: {
      deployment_ids: ids,
    },
    server: false,
  });
}

export async function DeleteDeploy(ids: string[], nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/deployments/deploy/delete`;

  return fetch(key, {
    method: 'POST',
    body: {
      deployment_ids: ids,
    },
    server: false,
  });
}
