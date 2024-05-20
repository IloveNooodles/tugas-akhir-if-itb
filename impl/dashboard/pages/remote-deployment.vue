<script setup lang="ts">
import { getDeploymentList } from '~/api/deployment';
import type { Deployment } from '~/types/deployment';
import { doRemoteDeploySchema } from '~/types/remote';

const nuxtApp = useNuxtApp();
const toast = useToast();

const {
  data: deployData,
  error: deployError,
  pending: deployPending,
  refresh: deployRefresh,
} = await getDeploymentList(nuxtApp);

const state = ref([]);
const deploymentList = ref<Deployment>();

const selectRepoOpts = computed(() => {
  return genereateSelectFromArray(deployData.value, 'name', 'id', []);
});

const isButtonDisabled = ref(false);

function addDeployments() {}


function deleteDeployments() {}
</script>

<template>
  <UContainer>
    <h1 class="text-center">Remote Deployment</h1>
    <UDivider />
    <h2>Do remote deployment</h2>
    <p>
      This will deploying the deployment from deployment available deployment
      list
    </p>
    <UCard v-if="deployPending"> Loading </UCard>
    <UCard v-else-if="deployError">Error</UCard>
    <UForm
      :schema="doRemoteDeploySchema"
      :state="state"
      class="space-y-4 pt-4"
      @submit="addDeployments"
      v-else
    >
      <UFormGroup label="Deployment list" name="repository_id">
        <USelectMenu
          v-model="state"
          :options="selectRepoOpts"
          option-attribute="name"
          multiple
          placeholder="Deployment"
        />
      </UFormGroup>

      <UButton type="submit" :disabled="isButtonDisabled"> Deploy </UButton>
    </UForm>
    <h2>Remove Deployment</h2>
  </UContainer>
</template>
