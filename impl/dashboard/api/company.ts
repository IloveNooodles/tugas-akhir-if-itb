import type { CompanyUser } from '~/types/company';
import type { Response } from '~/types/response';

function transformGetCompanyDetail(res: Response<CompanyUser>) {
  return res.data;
}

export async function getCompanyDetail(nuxtApp = useNuxtApp()) {
  const fetch = nuxtApp.$api;
  const key = `/api/v1/companies`;

  return useLazyFetch(key, {
    $fetch: fetch,
    transform: transformGetCompanyDetail,
    server: false,
  });
}
