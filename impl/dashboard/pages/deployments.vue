<script setup lang="ts">
const repoData = [
  {
    id: '12c234d2-bdb5-4725-b3d1-d275154e4c30',
    name: 'nginx',
    description: 'nginx latest',
    image: 'nginx:1.14.2',
    created_at: '2024-04-29T08:52:59.620887Z',
    updated_at: '2024-04-29T08:52:59.620887Z',
  },
];
const pending = false;
const columns = computed(() => {
  return generateColumnsFromArray(repoData, ['created_at', 'updated_at']);
});
</script>

<template>
  <UContainer>
    <h1 class="text-center">Deployments</h1>
    <UDivider />
    <div class="repository">
      <h2>Available Image</h2>
      <Table
        :data="repoData"
        :pending="pending"
        :columns="columns"
        :dropdown-items="dropdownItems"
      />
    </div>
  </UContainer>
</template>

<style scoped>
.repository {
  padding: 1rem;
  padding-left: 0;
}
</style>
