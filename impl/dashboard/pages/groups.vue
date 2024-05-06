<script setup lang="ts">
const groupData = [
  {
    id: 'b728b517-431f-4fac-8ba3-48d6672d82e1',
    name: 'home-deployment-group',
    company_id: '3a892d6b-817b-4913-b558-774a5e455c1c',
    created_at: '2024-04-22T14:39:20.818518Z',
    updated_at: '2024-04-22T14:39:20.818518Z',
  },
];
const pending = false;
const columns = computed(() => {
  return generateColumnsFromArray(groupData, [
    // 'created_at', 'updated_at'
  ]);
});

const dropdownList = dropdownItems;
</script>

<template>
  <UContainer>
    <h1 class="text-center">Groups</h1>
    <UDivider />
    <div class="wrap">
      <h2>Available Groups</h2>
      <Table
        :data="groupData"
        :pending="pending"
        :columns="columns"
        :dropdown-items="dropdownList"
      />
      <UButton
        label="Add Group"
        icon="i-heroicons-plus-solid"
        class="mt-2 max-w-fit self-end"
      />
    </div>
  </UContainer>
</template>
