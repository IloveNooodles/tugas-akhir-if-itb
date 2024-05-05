<script setup lang="ts">
import type { CompanyUser } from '~/types/company';

// const { data, status, pending, clear, execute, error } = await getUserLists();
// const {
//   data: companyData,
//   status: companyStatus,
//   pending: companyPending,
//   status,
// } = await getCompanyDetail();

const companyPending = false;
const companyData: CompanyUser = {
  id: '3a892d6b-817b-4913-b558-774a5e455c1c',
  name: 'gare-company',
  cluster_name: 'kind-cluster-multi-node',
  created_at: '2024-04-22T14:22:20.234039Z',
  updated_at: '2024-04-22T14:22:20.234039Z',
  email: 'mgarebaldhie@gmail.com',
  username: 'gawrgare',
};

const userPending = false;
const userData = [
  {
    id: '6e9ed8c2-5cec-4d94-a927-00cca14bf70e',
    name: 'gare-user-1',
    email: 'gare1@gmail.com',
    company_id: '3a892d6b-817b-4913-b558-774a5e455c1c',
    created_at: '2024-04-22T14:23:04.274123Z',
    updated_at: '2024-04-22T14:23:04.274123Z',
  },
  {
    id: '6e9ed8c2-5cec-4d94-a927-00cca14bf70f',
    name: 'gare-user-1',
    email: 'gare4@gmail.com',
    company_id: '3a892d6b-817b-4913-b558-774a5e455c1c',
    created_at: '2024-04-22T14:23:04.274123Z',
    updated_at: '2024-04-22T14:23:04.274123Z',
  },
  {
    id: '6e9ed8c2-5cec-4d94-a927-00cca14bf70g',
    name: 'gare-user-2',
    email: 'gare2@gmail.com',
    company_id: '3a892d6b-817b-4913-b558-774a5e455c1c',
    created_at: '2024-04-22T14:23:04.274123Z',
    updated_at: '2024-04-22T14:23:04.274123Z',
  },
  {
    id: '6e9ed8c2-5cec-4d94-a927-00cca14bf7he',
    name: 'gare-user-3',
    email: 'gare3@gmail.com',
    company_id: '3a892d6b-817b-4913-b558-774a5e455c1c',
    created_at: '2024-04-22T14:23:04.274123Z',
    updated_at: '2024-04-22T14:23:04.274123Z',
  },
];
</script>

<template>
  <div class="flex-1 flex flex-col p-4 gap-2">
    <h1 class="text-center">Account</h1>
    <UDivider label="Company" />
    <Company :company="companyData" :pending="companyPending" />
    <UDivider label="Users" />
    <UserList :users="userData" :pending="userPending" />
  </div>
</template>
