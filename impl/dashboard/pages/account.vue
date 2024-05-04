<script setup lang="ts">
import type { Company } from '~/types/company';

// const { data, status, pending, clear, execute, error } = await getUserLists();
// const {
//   data: companyData,
//   status: companyStatus,
//   pending: companyPending,
//   status,
// } = await getCompanyDetail();

const companyPending = true;
const companyData: Company = {
  id: '3a892d6b-817b-4913-b558-774a5e455c1c',
  name: 'gare-company',
  cluster_name: 'kind-cluster-multi-node',
  created_at: '2024-04-22T14:22:20.234039Z',
  updated_at: '2024-04-22T14:22:20.234039Z',
};
</script>

<template>
  <div class="flex-1 p-4">
    <h1 class="text-center">Account</h1>
    <UDivider label="Company" />
    <Company
      v-if="!companyPending"
      :key="companyData.id"
      :company="companyData"
    />
    <UDivider label="Users" />
    <UDivider label="List Users" />
  </div>
</template>
