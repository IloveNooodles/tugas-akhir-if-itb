<script setup lang="ts">
import { getCompanyDetail } from '~/api/company';
import { getUserLists } from '~/api/user';

const nuxtApp = useNuxtApp();

const {
  data: usersData,
  error: usersError,
  pending: usersPending,
} = await getUserLists(nuxtApp);

const {
  data: companyData,
  pending: companyPending,
  error: companyError,
} = await getCompanyDetail(nuxtApp);

const columns = computed(() => {
  return generateColumnsFromArray(usersData.value, [
    'created_at',
    'updated_at',
    'company_id',
  ]);
});
</script>

<template>
  <UContainer class="flex-1 flex flex-col gap-2">
    <h1 class="text-center">Account</h1>
    <UDivider />
    <div>
      <h2>Company</h2>
      <Company
        :error="companyError"
        :company="companyData"
        :pending="companyPending"
      />
    </div>
    <div>
      <h2>Users</h2>
      <UserList
        :data="usersData"
        :pending="usersPending"
        :columns="columns"
        :error="usersError"
      />
    </div>
  </UContainer>
</template>
