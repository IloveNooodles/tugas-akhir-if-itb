<script setup lang="ts">
import type { FormSubmitEvent } from '#ui/types';
import {
  createGroupDeviceSchema as schema,
  type CreateGroupDevice as Schema,
} from '@/types/groupdevice';
import { FetchError } from 'ofetch';
import { getDeploymentDetailByID } from '~/api/deployment';
import { createGroupDevice } from '~/api/groupdevice';
import { getDeploymentHistoryList } from '~/api/history';

const route = useRoute();
const nuxtApp = useNuxtApp();
const toast = useToast();
const deploymentId = route.params.id as string;

const {
  data: deploymentData,
  pending,
  error,
  refresh,
} = await getDeploymentDetailByID(deploymentId, nuxtApp);

// TODO rubah ini
const {
  data: deployHistoryData,
  error: deployHistoryError,
  pending: deployHistoryPending,
} = await getDeploymentHistoryList(nuxtApp);

const deployHistoryColumn = computed(() => {
  return generateColumnsFromArray(deployHistoryData.value, ['company_id']);
});

const isOpen = ref(false);
const disabled = ref(false);
const state = ref({
  group_id: '',
});

// TODO validation when creating group device
async function onSubmit(event: FormSubmitEvent<Schema>) {
  const body = event.data;
  try {
    await createGroupDevice(body, nuxtApp);
    toast.add({
      title: 'Success adding group',
    });

    disabled.value = true;
    isOpen.value = false;
  } catch (err: any) {
    if (err instanceof FetchError && err.data) {
      toast.add({ title: err.data.message, color: 'red' });
      return;
    }

    toast.add({ title: 'Please try again', color: 'red' });
  } finally {
    disabled.value = false;
    state.value.group_id = '';
  }
}
</script>

<template>
  <Loading v-if="pending" class="flex justify-center items-center h-full" />
  <p v-else-if="error" class="flex justify-center items-center h-full">
    Server error, please try again
  </p>
  <UContainer v-else>
    <div>
      <div class="flex flex-row items-center justify-between">
        <h1 class="m-0">{{ deploymentData.name }}</h1>
        <div class="btnContainer flex gap-2">
          <UButton icon="i-heroicons-pencil-square-20-solid" />
          <UButton icon="i-heroicons-trash-20-solid" />
        </div>
      </div>
      <UDivider />
      <div class="flex flex-col gap-2 p-4 pl-0">
        <div class="wrap">
          <UCard>
            <h2>Description</h2>
            <p>Node name: {{ deploymentData.name }}</p>
          </UCard>
        </div>
        <div class="flex flex-col">
          <UModal v-model="isOpen">
            <UCard>
              <h2 class="text-center p-0 m-0">Add New Group</h2>
              <UForm
                :schema="schema"
                :state="state"
                class="space-y-4 pt-4"
                @submit="onSubmit"
              >
              </UForm>
            </UCard>
          </UModal>
        </div>
        <div>
          <h2>Histories</h2>
          <HistoryList
            :data="deployHistoryData"
            :pending="deployHistoryPending"
            :columns="deployHistoryColumn"
            :error="deployHistoryError"
          />
        </div>
      </div>
    </div>
  </UContainer>
</template>

<style scoped>
h1,
h2 {
  padding: 0;
}
.wrap {
  width: fit-content;
}
</style>
