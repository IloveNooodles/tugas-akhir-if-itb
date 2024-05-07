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

const dropdownItems = computed(() => {
  return (row: any) => [
    [
      {
        label: 'Detail',
        icon: 'i-heroicons-document-magnifying-glass-16-solid',
        click: () => {
          navigateTo(`/groups/${row.id}`);
        },
      },
      {
        label: 'Edit',
        icon: 'i-heroicons-pencil-square-20-solid',
      },
    ],
    [
      {
        label: 'Delete',
        icon: 'i-heroicons-trash-20-solid',
      },
    ],
  ];
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
      />
    </div>
  </UContainer>
</template>
