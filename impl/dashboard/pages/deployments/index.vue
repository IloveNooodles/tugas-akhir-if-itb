<script setup lang="ts">
import type { FormSubmitEvent } from '#ui/types';
import { FetchError } from 'ofetch';
import { createDeployment, getDeploymentList } from '~/api/deployment';
import { createRepository, getRepositoryList } from '~/api/repository';
import {
  type CreateDeploymentSchema,
  createDeploymentSchema,
} from '~/types/deployment';
import {
  type CreateRepositorySchema,
  createRepositorySchema,
} from '~/types/repository';

const nuxtApp = useNuxtApp();
const toast = useToast();

const {
  data: repoData,
  error: repoError,
  pending: repoPending,
  refresh: repoRefresh,
} = await getRepositoryList(nuxtApp);

const repoColumns = computed(() => {
  return generateColumnsFromArray(repoData.value, []);
});

const selectRepoOpts = computed(() => {
  return genereateSelectFromArray(repoData.value, 'name', 'id', []);
});

const {
  data: deployData,
  error: deployError,
  pending: deployPending,
  refresh: deployRefresh,
} = await getDeploymentList(nuxtApp);

// TODO gabungin sama image biar jelas tablenya ngapain
const deployColumn = computed(() => {
  return generateColumnsFromArray(deployData.value, []);
});

const repositoryState = ref<CreateRepositorySchema>({
  description: '',
  image: '',
  name: '',
});

const deploymentState = ref<CreateDeploymentSchema>({
  name: '',
  repository_id: '',
  target: '',
  version: '',
});

const isModalRepoDisabled = ref(false);
const isButtonRepoDisabled = ref(false);

const isModalDeployDisabled = ref(false);
const isButtonDeployDisabled = ref(false);

async function addRepository(event: FormSubmitEvent<CreateRepositorySchema>) {
  const body = event.data;
  try {
    await createRepository(body, nuxtApp);
    toast.add({
      title: `Success creating repository ${body.name}`,
    });

    isButtonRepoDisabled.value = true;
    isModalRepoDisabled.value = true;
    await repoRefresh();
  } catch (err: any) {
    if (err instanceof FetchError && err.data) {
      toast.add({ title: err.data.message, color: 'red' });
      return;
    }

    toast.add({ title: 'Please try again', color: 'red' });
  } finally {
    isButtonRepoDisabled.value = false;
    isModalRepoDisabled.value = false;

    repositoryState.value.name = '';
    repositoryState.value.description = '';
    repositoryState.value.image = '';
  }
}

async function addDeployments(event: FormSubmitEvent<CreateDeploymentSchema>) {
  const body = event.data;
  try {
    await createDeployment(body, nuxtApp);
    toast.add({
      title: `Success creating deployment ${body.name}`,
    });

    isButtonDeployDisabled.value = true;
    isModalDeployDisabled.value = true;
    await deployRefresh();
  } catch (err: any) {
    if (err instanceof FetchError && err.data) {
      toast.add({ title: err.data.message, color: 'red' });
      return;
    }

    toast.add({ title: 'Please try again', color: 'red' });
  } finally {
    isButtonDeployDisabled.value = false;
    isModalDeployDisabled.value = false;

    repositoryState.value.name = '';
    repositoryState.value.description = '';
    repositoryState.value.image = '';
  }
}

async function onDeleteRepository() {
  await repoRefresh();
}

async function onDeleteDeployments() {
  await deployRefresh();
}
</script>

<template>
  <UContainer>
    <h1 class="text-center">Deployments</h1>
    <UDivider />
    <div class="wrap">
      <h2>Deployments</h2>
      <DeploymentList
        :data="deployData"
        :pending="deployPending"
        :columns="deployColumn"
        :error="deployError"
        @on-delete="onDeleteDeployments"
      />
      <UButton
        label="Add Deployment"
        icon="i-heroicons-plus-solid"
        class="mt-2 max-w-fit self-end"
        @click="isModalDeployDisabled = !isModalDeployDisabled"
      />
      <UModal v-model="isModalDeployDisabled">
        <UCard>
          <h2 class="text-center p-0 m-0">Add New Deployment</h2>
          <UForm
            :schema="createDeploymentSchema"
            :state="deploymentState"
            class="space-y-4 pt-4"
            @submit="addDeployments"
          >
            <UFormGroup label="Name" name="name">
              <UInput v-model="deploymentState.name" type="text" />
            </UFormGroup>
            <UFormGroup label="Version" name="version">
              <UInput v-model="deploymentState.version" type="text" />
            </UFormGroup>
            <UFormGroup label="Target" name="target">
              <UInput v-model="deploymentState.target" type="text" />
            </UFormGroup>
            <UFormGroup label="Repository" name="repository_id">
              <USelect
                v-model="deploymentState.repository_id"
                :options="selectRepoOpts"
                option-attribute="name"
                placeholder="repository"
              />
            </UFormGroup>

            <UButton type="submit" :disabled="isButtonDeployDisabled">
              Submit
            </UButton>
          </UForm>
        </UCard>
      </UModal>
    </div>
    <div class="wrap">
      <h2>Images</h2>
      <RepositoryList
        :data="repoData"
        :pending="repoPending"
        :error="repoError"
        :columns="repoColumns"
        @on-delete="onDeleteRepository"
      />
      <UButton
        label="Add image"
        icon="i-heroicons-plus-solid"
        class="mt-2 max-w-fit self-end"
        @click="isModalRepoDisabled = !isModalRepoDisabled"
      />
      <UModal v-model="isModalRepoDisabled">
        <UCard>
          <h2 class="text-center p-0 m-0">Add New Repository</h2>
          <UForm
            :schema="createRepositorySchema"
            :state="repositoryState"
            class="space-y-4 pt-4"
            @submit="addRepository"
          >
            <UFormGroup label="Name" name="name">
              <UInput v-model="repositoryState.name" type="text" />
            </UFormGroup>
            <UFormGroup label="Description" name="description">
              <UInput v-model="repositoryState.description" type="text" />
            </UFormGroup>
            <UFormGroup label="Image" name="image">
              <UInput v-model="repositoryState.image" type="text" />
            </UFormGroup>

            <UButton type="submit" :disabled="isButtonRepoDisabled">
              Submit
            </UButton>
          </UForm>
        </UCard>
      </UModal>
    </div>

    <!-- <div class="wrap">
      <h2>Deployment Histories</h2>
      <HistoryList
        :data="deployHistoryData"
        :pending="deployHistoryPending"
        :columns="deployHistoryColumn"
        :error="deployHistoryError"
      />
    </div> -->
  </UContainer>
</template>
