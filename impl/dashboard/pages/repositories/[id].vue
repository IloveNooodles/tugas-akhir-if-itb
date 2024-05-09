<script setup lang="ts">
import { FetchError } from 'ofetch';
import {
  deleteDeploymentByID,
  getDeploymentDetailByID,
} from '~/api/deployment';
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
// TODO gabung sama data image juga make yg mana jadi di join
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

const isConfirmModal = ref(false);
const isButtonDisabled = ref(false);

async function deleteDevices() {
  try {
    isButtonDisabled.value = true;
    await deleteDeploymentByID(deploymentId, nuxtApp);
    toast.add({ title: `Success deleteing deployment ${deploymentId}` });
    await navigateTo('/deployments');
  } catch (err) {
    if (err instanceof FetchError && err.data) {
      toast.add({ title: err.data.message, color: 'red' });
      return;
    }

    toast.add({ title: 'Please try again', color: 'red' });
  } finally {
    isButtonDisabled.value = false;
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
          <UButton
            icon="i-heroicons-trash-20-solid"
            :disabled="isButtonDisabled"
            @click="isConfirmModal = !isConfirmModal"
          />
          <UModal v-model="isConfirmModal">
            <UCard>
              <h2 class="text-center p-0 m-0">
                Are you sure you want to delete
              </h2>
              <h3 class="text-center">{{ deploymentData.name }}</h3>
              <div class="flex justify-center gap-10">
                <UButton
                  type="submit"
                  :disabled="isButtonDisabled"
                  @click.prevent="deleteDevices"
                >
                  Yes
                </UButton>
                <UButton
                  type="submit"
                  @click.prevent="isConfirmModal = false"
                  color="red"
                >
                  No
                </UButton>
              </div>
            </UCard>
          </UModal>
        </div>
      </div>
      <UDivider />
      <div class="flex flex-col gap-2 p-4 pl-0">
        <div class="wrap">
          <UCard>
            <h2>Description</h2>
            <p>name: {{ deploymentData.name }}</p>
          </UCard>
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
