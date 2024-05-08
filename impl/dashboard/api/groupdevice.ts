import type { CreateGroupDevice } from '~/types/groupdevice';

export async function createGroupDevice(
  body: CreateGroupDevice,
  nuxtApp = useNuxtApp(),
) {
  const fetch = nuxtApp.$api;
  const key = '/api/v1/groupdevices';
  return await fetch(key, {
    method: 'POST',
    body: body,
    timeout: 500,
  });
}

export async function deleteGroupDeviceByID(
  id: string,
  nuxtApp = useNuxtApp(),
) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/groupdevices/${id}`;

  return fetch(key, {
    method: 'DELETE',
    timeout: 500,
  });
}
