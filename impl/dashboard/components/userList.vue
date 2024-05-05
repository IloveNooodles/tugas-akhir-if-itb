<script setup lang="ts">
import type { User } from '~/types/user';

interface Props {
  users: User[];
  pending: boolean;
}

const props = defineProps<Props>();
const { users, pending } = props;
// TODO remove unrelated fields like created_at, updated_at, id
const selected = ref([]);
const dropdownItems = (row) => [
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
];
const columns = generateColumnsFromArray(users);
</script>

<template>
  <div>
    <UTable
      :rows="users"
      :columns="columns"
      v-model="selected"
      :loading="false"
      class="bg-slate-900 rounded-lg"
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
      <template #loading-state>
        <div class="flex items-center justify-center h-32">
          <i class="loader --6" />
        </div>
      </template>
      <template #empty-state>
        <div class="flex flex-col items-center justify-center py-6 gap-3">
          <span class="italic text-sm">No one here!</span>
          <UButton label="Add people" />
        </div>
      </template>
    </UTable>
  </div>
</template>
