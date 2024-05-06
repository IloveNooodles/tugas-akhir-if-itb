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
</script>

<template>
  <div class="flex-1 flex flex-col p-4 gap-2">
    <h1 class="text-center">Account</h1>
    <UDivider label="Company" />
    <Company
      :error="companyError"
      :company="companyData"
      :pending="companyPending"
    />
    <UDivider label="Users" />
    <UserList :users="usersData" :pending="usersPending" :error="usersError" />
  </div>
</template>
