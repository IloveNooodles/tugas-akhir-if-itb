import type { Group } from '~/types/group';
import type { Response } from '~/types/response';

function transformGetGroupList(res: Response<Array<Group>>) {
  return res.data;
}

export async function getGroupList(nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/groups`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetGroupList,
    server: false,
  });
}
