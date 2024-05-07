<script setup lang="ts">
import type { User } from '~/types/user';

interface Props {
  data: User[] | null;
  pending: boolean;
  columns: {
    [key: string]: any;
    key: string;
    sortable?: boolean | undefined;
    sort?: ((a: any, b: any, direction: 'asc' | 'desc') => number) | undefined;
    direction?: 'asc' | 'desc' | undefined;
    class?: string | undefined;
  }[];
  error?: any;
}

const props = defineProps<Props>();
const dropdownItems = computed(() => {
  return (row: any) => [
    [
      {
        label: 'Detail',
        icon: 'i-heroicons-document-magnifying-glass-16-solid',
        click: () => {
          navigateTo(`/account/${row.id}`);
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
const columns = computed(() => {
  return generateColumnsFromArray(props.data, [
    'created_at',
    'updated_at',
    'company_id',
  ]);
});
</script>

<template>
  <UCard v-if="pending">
    <Loading />
  </UCard>
  <UCard v-else-if="error || !data">
    <div>Sorry, we're having an issue please try again</div>
  </UCard>
  <Table
    v-else
    :data="data"
    :pending="pending"
    :columns="columns"
    :dropdown-items="dropdownItems"
  >
  </Table>
</template>
