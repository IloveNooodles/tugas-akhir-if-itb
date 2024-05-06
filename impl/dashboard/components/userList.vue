<script setup lang="ts">
import type { User } from '~/types/user';

interface Props {
  users: User[] | null;
  pending: boolean;
  error?: any;
}

const props = defineProps<Props>();
const dropdownItems = computed(() => (row) => [
  [
    {
      label: 'Edit',
      icon: 'i-heroicons-pencil-square-20-solid',
      click: () => console.log('Edit', row.id),
    },
    {
      label: 'Duplicate',
      icon: 'i-heroicons-document-duplicate-20-solid',
    },
  ],
  [
    {
      label: 'Archive',
      icon: 'i-heroicons-archive-box-20-solid',
    },
    {
      label: 'Move',
      icon: 'i-heroicons-arrow-right-circle-20-solid',
    },
  ],
  [
    {
      label: 'Delete',
      icon: 'i-heroicons-trash-20-solid',
    },
  ],
]);
const columns = computed(() => {
  return generateColumnsFromArray(props.users, [
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
  <UCard v-else-if="error || !users">
    <div>Sorry, we're having an issue please try again</div>
  </UCard>
  <UTable
    :rows="users"
    :columns="columns"
    :loading="false"
    class="bg-slate-900 rounded-lg"
    v-else
  >
    <template #actions-data="{ row }">
      <UDropdown :items="dropdownItems(row)">
        <UButton
          color="gray"
          variant="ghost"
          icon="i-heroicons-ellipsis-horizontal-20-solid"
        />
      </UDropdown>
    </template>
    <template #empty-state>
      <div class="flex flex-col items-center justify-center py-6 gap-3">
        <span class="italic text-sm">No one here!</span>
        <UButton label="Add people" />
      </div>
    </template>
  </UTable>
</template>
