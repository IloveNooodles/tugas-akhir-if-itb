import type { Device } from '~/types/device';
import type { Response } from '~/types/response';

function transformGetDeviceList(res: Response<Array<Device>>) {
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
