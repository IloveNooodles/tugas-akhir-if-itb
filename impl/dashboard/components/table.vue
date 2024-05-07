<script setup lang="ts" generic="T">
interface Props {
  data: T[] | null;
  pending: boolean;
  columns: {
    [key: string]: any;
    key: string;
    sortable?: boolean | undefined;
    sort?: ((a: any, b: any, direction: 'asc' | 'desc') => number) | undefined;
    direction?: 'asc' | 'desc' | undefined;
    class?: string | undefined;
  }[];
  dropdownItems: GenerateDropdownItems;
  error?: any;
}

const props = defineProps<Props>();
</script>

<template>
  <UCard v-if="pending">
    <Loading />
  </UCard>
  <UCard v-else-if="error || !data">
    <div>Sorry, we're having an issue please try again</div>
  </UCard>
  <UTable
    :rows="data"
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
        <span class="italic text-sm">It's empty</span>
        <UButton label="Add Items" />
      </div>
    </template>
  </UTable>
</template>
