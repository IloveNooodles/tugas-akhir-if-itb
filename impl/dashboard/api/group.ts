import type { CreateGroupSchema, Group } from '~/types/group';
import type { GroupDeviceG } from '~/types/groupdevice';
import type { Response } from '~/types/response';

function transformGetGroupList(res: Response<Array<Group>>) {
  return res.data;
}

function transformGetGroupById(res: Response<Group>) {
  return res.data;
}

function transformGetGroupDeviceById(res: Response<Array<GroupDeviceG>>) {
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

export async function getGroupById(id: string, nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/groups/${id}`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetGroupById,
    server: false,
  });
}

export async function getGroupDevicesById(id: string, nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/groups/${id}/devices`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetGroupDeviceById,
    server: false,
  });
}

export async function createGroup(
  body: CreateGroupSchema,
  nuxtApp = useNuxtApp(),
) {
  const fetch = nuxtApp.$api;
  const key = '/api/v1/groups';
  return await fetch(key, {
    method: 'POST',
    body: body,
    timeout: 500,
  });
}

export async function deleteGroupByID(id: string, nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/groups/${id}`;

  return fetch(key, {
    method: 'DELETE',
    timeout: 500,
  });
}