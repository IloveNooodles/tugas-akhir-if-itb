<script setup lang="ts">
import type { FormSubmitEvent } from '#ui/types';
import {
  createGroupDeviceSchema as schema,
  type CreateGroupDevice as Schema,
} from '@/types/groupdevice';
import { FetchError } from 'ofetch';
import {
  deleteDeviceByID,
  getDeviceById,
  getDeviceGroupsById,
} from '~/api/device';
import { getGroupList } from '~/api/group';
import { createGroupDevice } from '~/api/groupdevice';
import { getDeploymentHistoryListWithDevices } from '~/api/history';

const route = useRoute();
const nuxtApp = useNuxtApp();
const toast = useToast();
const deviceId = route.params.id as string;

const {
  data: deviceData,
  pending,
  error,
  refresh,
} = await getDeviceById(deviceId, nuxtApp);

const {
  data: groupData,
  error: groupError,
  pending: groupPending,
  refresh: groupRefresh,
} = await getDeviceGroupsById(deviceId, nuxtApp);

const {
  data: deployHistoryData,
  error: deployHistoryError,
  pending: deployHistoryPending,
} = await getDeploymentHistoryListWithDevices(nuxtApp, deviceId);

const { data: groupListData } = await getGroupList(nuxtApp);

const columns = computed(() => {
  return generateColumnsFromArray(groupData.value, ['company_id']);
});

const deployHistoryColumn = computed(() => {
  return generateColumnsFromArray(deployHistoryData.value, ['company_id']);
});

const isOpen = ref(false);
const disabled = ref(false);
const state = ref({
  group_id: '',
  device_id: deviceId,
});

const isConfirmModal = ref(false);
const isButtonDisabled = ref(false);

const listSelectGroup = computed(() => {
  return genereateSelectFromArray(
    groupListData.value,
    'name',
    'id',
    groupData.value,
  );
});

async function onSubmit(event: FormSubmitEvent<Schema>) {
  const body = event.data;
  try {
    await createGroupDevice(body, nuxtApp);
    toast.add({
      title: 'Success adding group',
    });

    disabled.value = true;
    isOpen.value = false;
    await groupRefresh();
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

async function deleteDevices() {
  try {
    isButtonDisabled.value = true;
    await deleteDeviceByID(deviceId, nuxtApp);
    toast.add({ title: `Success Deleteing device ${deviceId}` });
    await navigateTo('/devices');
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

function onDelete() {
  groupRefresh();
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
        <h1 class="m-0">{{ deviceData.name }}</h1>
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
              <h3 class="text-center">{{ deviceData.name }}</h3>
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
            <p>Node name: {{ deviceData.node_name }}</p>
          </UCard>
        </div>
        <div class="flex flex-col">
          <h2>Groups</h2>
          <GroupdeviceList
            :data="groupData"
            :pending="groupPending"
            :columns="columns"
            :error="groupError"
            @on-delete="onDelete"
          />
          <UButton
            label="Add Group"
            icon="i-heroicons-plus-solid"
            class="mt-2 max-w-fit self-end"
            @click="isOpen = !isOpen"
          />
          <UModal v-model="isOpen">
            <UCard>
              <h2 class="text-center p-0 m-0">Add New Group</h2>
              <UForm
                :schema="schema"
                :state="state"
                class="space-y-4 pt-4"
                @submit="onSubmit"
              >
                <UFormGroup label="Group" name="group_id">
                  <USelect
                    v-model="state.group_id"
                    :options="listSelectGroup"
                    option-attribute="name"
                    placeholder="Groups"
                  />
                </UFormGroup>
                <UButton type="submit" :disabled="disabled"> Submit </UButton>
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
