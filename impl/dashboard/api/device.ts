import type { CreateDeviceSchema, Device } from '~/types/device';
import type { Group } from '~/types/group';
import type { Response } from '~/types/response';

function transformGetDeviceList(res: Response<Array<Device>>) {
  return res.data;
}

function transformGetDevice(res: Response<Device>) {
  return res.data;
}

function transformGetDeviceGroups(res: Response<Array<Group>>) {
  return res.data;
}

export async function getDeviceList(nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/devices`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetDeviceList,
    server: false,
  });
}

export async function getDeviceById(id: string, nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/devices/${id}`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetDevice,
    server: false,
  });
}

export async function getDeviceGroupsById(id: string, nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/devices/${id}/groups`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetDeviceGroups,
    server: false,
  });
}

export async function createDevice(
  body: CreateDeviceSchema,
  nuxtApp = useNuxtApp(),
) {
  const fetch = nuxtApp.$api;
  const key = '/api/v1/devices';
  return await fetch(key, {
    method: 'POST',
    body: body,
    timeout: 500,
  });
}
