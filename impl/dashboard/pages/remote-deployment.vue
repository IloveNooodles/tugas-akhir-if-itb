<script setup lang="ts">
import type { FormSubmitEvent } from '#ui/types';
import { FetchError } from 'ofetch';
import { getDeploymentList } from '~/api/deployment';
import { DeleteDeploy, Deploy } from '~/api/remote';
import {
  doRemoteDeploySchema,
  type DoRemoteDeploySchema,
} from '~/types/remote';

const nuxtApp = useNuxtApp();
const toast = useToast();

const {
  data: deployData,
  error: deployError,
  pending: deployPending,
  refresh: deployRefresh,
} = await getDeploymentList(nuxtApp);

const stateDeleteDeployment = ref([]);
const stateDeployment = ref({
  deployment_ids: [],
  type: 'TARGET',
});

const selectedTypeDeployment = ref('TARGET');
const selectedDeploymentIds = ref([]);

const selectRepoOpts = computed(() => {
  return genereateSelectFromArray(deployData.value, 'name', 'id', []);
});

const isButtonDeploymentDisabled = ref(false);
const isButtonDeleteDeploymentDisabled = ref(false);
const availableTypeDeploy = ['TARGET', 'CUSTOM'];

async function addDeployments(event: FormSubmitEvent<DoRemoteDeploySchema>) {
  const body = event.data;
  body.deployment_ids = selectedDeploymentIds.value.map((x) => x.value);
  try {
    const response = await Deploy(body, nuxtApp);
    toast.add({
      title: `Success doing remote deployment`,
    });
    isButtonDeploymentDisabled.value = true;
  } catch (err: any) {
    if (err instanceof FetchError && err.data) {
      toast.add({ title: err.data.message, color: 'red' });
      return;
    }

    toast.add({ title: 'Please try again', color: 'red' });
  } finally {
    isButtonDeploymentDisabled.value = false;
    stateDeployment.value.deployment_ids = [];
    stateDeployment.value.type = 'TARGET';
  }
}

async function deleteDeployments(event) {
  const deployment_ids = stateDeleteDeployment.value.map((x) => x.value);

  try {
    const response = await DeleteDeploy(deployment_ids, nuxtApp);
    toast.add({
      title: `Success deleting remote deployment`,
    });
    isButtonDeploymentDisabled.value = true;
  } catch (err: any) {
    if (err instanceof FetchError && err.data) {
      toast.add({ title: err.data.message, color: 'red' });
      return;
    }

    toast.add({ title: 'Please try again', color: 'red' });
  } finally {
    isButtonDeploymentDisabled.value = false;
    stateDeployment.value.deployment_ids = [];
    stateDeployment.value.type = 'TARGET';
  }
}

watch(stateDeleteDeployment, () => {
  console.log(stateDeleteDeployment);
});
</script>

<template>
  <UContainer>
    <h1 class="text-center">Remote Deployment</h1>
    <UDivider />
    <div>
      <h2 class="font-bold">Do remote deployment</h2>
      <p>Select deployment that you want to deploy</p>
      <UCard v-if="deployPending"> Loading </UCard>
      <UCard v-else-if="deployError">Error</UCard>
      <UForm
        :schema="doRemoteDeploySchema"
        :state="stateDeployment"
        class="space-y-4 pt-4 mb-8"
        @submit="addDeployments"
        v-else
      >
        <UFormGroup label="Deployment list" name="deployment_ids">
          <USelectMenu
            v-model="selectedDeploymentIds"
            :options="selectRepoOpts"
            option-attribute="name"
            multiple
            placeholder="Available Deployments"
          />
        </UFormGroup>

        <UFormGroup label="Type" name="type">
          <USelectMenu
            v-model="stateDeployment.type"
            :options="availableTypeDeploy"
            placeholder="Type"
          />
        </UFormGroup>
        <UButton type="submit" :disabled="isButtonDeploymentDisabled">
          Deploy
        </UButton>
      </UForm>
    </div>
    <UDivider class="m-b-4" />
    <div>
      <h2 class="font-bold">Remove Deployment</h2>
      <p>Select deployment that you want to remove</p>
      <UCard v-if="deployPending"> Loading </UCard>
      <UCard v-else-if="deployError">Error</UCard>
      <UForm
        :state="stateDeleteDeployment"
        class="space-y-4 pt-4 mb-8"
        @submit="deleteDeployments"
        v-else
      >
        <UFormGroup label="Deployment list" name="deployment_ids">
          <USelectMenu
            v-model="stateDeleteDeployment"
            :options="selectRepoOpts"
            option-attribute="name"
            multiple
            placeholder="Available Deployments"
          />
        </UFormGroup>
        <UButton type="submit" :disabled="isButtonDeleteDeploymentDisabled">
          Remove
        </UButton>
      </UForm>
    </div>
  </UContainer>
</template>
