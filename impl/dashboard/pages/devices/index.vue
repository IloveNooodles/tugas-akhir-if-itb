<script setup lang="ts">
import type { FormSubmitEvent } from '#ui/types';
import {
  createDeviceSchema as schema,
  type CreateDeviceSchema as Schema,
} from '@/types/device';
import { FetchError } from 'ofetch';
import { createDevice, getDeviceList } from '~/api/device';

const nuxtApp = useNuxtApp();
const toast = useToast();

const {
  data: devicesData,
  error: devicesError,
  pending: devicesPending,
  refresh: devicesRefresh,
} = await getDeviceList(nuxtApp);

const columns = computed(() => {
  return generateColumnsFromArray(devicesData.value, [
    'created_at',
    'updated_at',
  ]);
});

const isOpen = ref(false);
const disabled = ref(false);
const state = ref({
  name: '',
  attributes: '',
  node_name: '',
  type: '',
});

async function onSubmit(event: FormSubmitEvent<Schema>) {
  const body = event.data;
  console.log(body);
  try {
    await createDevice(body, nuxtApp);
    toast.add({
      title: 'Success Creating Device',
    });

    disabled.value = true;
    isOpen.value = false;

    await devicesRefresh();
  } catch (err: any) {
    disabled.value = false;

    if (err instanceof FetchError && err.data) {
      toast.add({ title: err.data.message, color: 'red' });
      return;
    }

    toast.add({ title: 'Please try again', color: 'red' });
  }
}

async function onDelete() {
  await devicesRefresh();
}
</script>

<template>
  <UContainer>
    <h1 class="text-center">Devices</h1>
    <UDivider />
    <div class="wrap">
      <h2>Available Devices</h2>
      <DeviceList
        :data="devicesData"
        :pending="devicesPending"
        :columns="columns"
        :error="devicesError"
        @on-delete="onDelete"
      />
      <UButton
        label="Add Device"
        icon="i-heroicons-plus-solid"
        class="mt-2 max-w-fit self-end"
        @click="isOpen = !isOpen"
      />
      <UModal v-model="isOpen">
        <UCard>
          <h2 class="text-center p-0 m-0">Add New Device</h2>
          <UForm
            :schema="schema"
            :state="state"
            class="space-y-4 pt-4"
            @submit="onSubmit"
          >
            <UFormGroup label="Name" name="name">
              <UInput v-model="state.name" type="text" />
            </UFormGroup>

            <UFormGroup label="Node Name" name="node_name">
              <UInput v-model="state.node_name" type="text" />
            </UFormGroup>

            <UFormGroup label="Type" name="type">
              <UInput v-model="state.type" type="text" />
            </UFormGroup>

            <UFormGroup label="Attributes" name="attributes">
              <UInput v-model="state.attributes" type="text" />
            </UFormGroup>
            <UButton type="submit" :disabled="disabled"> Submit </UButton>
          </UForm>
        </UCard>
      </UModal>
    </div>
  </UContainer>
</template>
