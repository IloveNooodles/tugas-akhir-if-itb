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
  return generateColumnsFromArray(repoData, [
    // 'created_at', 'updated_at'
  ]);
});

// TODO gabungin sama image biar jelas tablenya ngapain
// TODO Tambahkan create juga di buttonlistnya
const deployData = [
  {
    id: '429bb8de-da58-4434-a5e0-3c3610f630c3',
    repository_id: '12c234d2-bdb5-4725-b3d1-d275154e4c30',
    name: 'home-deployment-group',
    version: 'v1',
    target: 'app=raspi',
    created_at: '2024-04-29T08:53:34.619907Z',
    updated_at: '2024-04-29T08:53:34.619907Z',
  },
];
const deployPending = false;
const deployColumn = computed(() => {
  return generateColumnsFromArray(deployData, [
    // 'created_at',
    // 'updated_at',
    // 'repository_id',
  ]);
});

const deployHistoryData = [
  {
    id: '10da31f8-9df4-48e9-8698-052599d49415',
    device_id: 'c4b21c55-8520-4aab-a789-75ad6ed88dc2',
    image_id: '12c234d2-bdb5-4725-b3d1-d275154e4c30',
    deployment_id: '429bb8de-da58-4434-a5e0-3c3610f630c3',
    status: 'DEPLOYING',
    created_at: '2024-04-29T08:59:19.327526Z',
    updated_at: '2024-04-29T08:59:19.327526Z',
  },
];
const deployHistoryPending = false;
const deployHistoryColumn = computed(() => {
  return generateColumnsFromArray(deployHistoryData, [
    // 'created_at',
    // 'updated_at',
    // 'repository_id',
  ]);
});
</script>

<template>
  <UContainer>
    <h1 class="text-center">Deployments</h1>
    <UDivider />
    <div class="wrap">
      <h2>Images</h2>
      <RepositoryList :data="repoData" :pending="pending" :columns="columns" />
      <UButton
        label="Add image"
        icon="i-heroicons-plus-solid"
        class="mt-2 max-w-fit self-end"
      />
    </div>
    <div class="wrap">
      <h2>Deployments</h2>
      <DeploymentList
        :data="deployData"
        :pending="deployPending"
        :columns="deployColumn"
      />
      <UButton
        label="Add Deployment"
        icon="i-heroicons-plus-solid"
        class="mt-2 max-w-fit self-end"
      />
    </div>
    <div class="wrap">
      <h2>Deployment Histories</h2>
      <HistoryList
        :data="deployHistoryData"
        :pending="deployHistoryPending"
        :columns="deployHistoryColumn"
      />
    </div>
  </UContainer>
</template>
