<script setup lang="ts">
import { getDeploymentList } from '~/api/deployment';
import { getDeploymentHistoryList } from '~/api/history';
import { getRepositoryList } from '~/api/repository';

const nuxtApp = useNuxtApp();

const {
  data: repoData,
  error: repoError,
  pending: repoPending,
} = await getRepositoryList(nuxtApp);


const repoColumns = computed(() => {
  return generateColumnsFromArray(repoData.value, [
    // 'created_at', 'updated_at'
  ]);
});

const {
  data: deployData,
  error: deployError,
  pending: deployPending,
} = await getDeploymentList(nuxtApp);

// TODO gabungin sama image biar jelas tablenya ngapain
const deployColumn = computed(() => {
  return generateColumnsFromArray(deployData.value, [
    // 'created_at',
    // 'updated_at',
    // 'repository_id',
  ]);
});

const {
  data: deployHistoryData,
  error: deployHistoryError,
  pending: deployHistoryPending,
} = await getDeploymentHistoryList(nuxtApp);

const deployHistoryColumn = computed(() => {
  return generateColumnsFromArray(deployHistoryData.value, [
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
      <RepositoryList
        :data="repoData"
        :pending="repoPending"
        :error="repoError"
        :columns="repoColumns"
      />
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
        :error="deployError"
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
