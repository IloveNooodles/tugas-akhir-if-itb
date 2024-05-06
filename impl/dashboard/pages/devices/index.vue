<script setup lang="ts">
const devicesData = [
  {
    id: 'c4b21c55-8520-4aab-a789-75ad6ed88dc2',
    name: 'raspberrypi-3',
    type: 'device',
    company_id: '3a892d6b-817b-4913-b558-774a5e455c1c',
    attributes: 'fire',
    node_name: 'cluster-multi-node-worker',
    created_at: '2024-04-22T14:39:02.474357Z',
    updated_at: '2024-04-22T14:39:02.474357Z',
  },
];
const pending = false;
const columns = computed(() => {
  return generateColumnsFromArray(devicesData, [
    // 'created_at', 'updated_at'
  ]);
});

const dropdownList = dropdownItems;
</script>

<template>
  <UContainer>
    <h1 class="text-center">Devices</h1>
    <UDivider />
    <div class="wrap">
      <h2>Available Devices</h2>
      <Table
        :data="devicesData"
        :pending="pending"
        :columns="columns"
        :dropdown-items="dropdownList"
      />
      <UButton
        label="Add Device"
        icon="i-heroicons-plus-solid"
        class="mt-2 max-w-fit self-end"
      />
    </div>
  </UContainer>
</template>
