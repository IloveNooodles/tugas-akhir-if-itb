<script setup lang="ts">
import { getCompanyDetail } from '~/api/company';
import { getUserLists } from '~/api/user';
import Table from '~/components/table.vue';

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

const dropdownList = dropdownItems;
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
      <Table
        :data="usersData"
        :pending="usersPending"
        :columns="columns"
        :dropdown-items="dropdownList"
        :error="usersError"
      />
    </div>
  </UContainer>
</template>
