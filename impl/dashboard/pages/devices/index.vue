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
  try {
    await createDevice(body, nuxtApp);
    toast.add({
      title: 'Success Creating Device',
    });
    disabled.value = true;
    isOpen.value = false;
    devicesRefresh();
  } catch (err: any) {
    disabled.value = false;
    if (err instanceof FetchError && err.data) {
      toast.add({ title: err.data.message, color: 'red' });
      return;
    }

    toast.add({ title: 'Please try again', color: 'red' });
  }
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
      />
      <UButton
        label="Add Device"
        icon="i-heroicons-plus-solid"
        class="mt-2 max-w-fit self-end"
        @click="isOpen = !isOpen"
      />
      <UModal v-model="isOpen">
        <UCard>
          <UForm
            :schema="schema"
            :state="state"
            class="space-y-4 pt-4"
            @submit="onSubmit"
          >
            <UFormGroup label="Name" name="Name">
              <UInput v-model="state.name" />
            </UFormGroup>

            <UFormGroup label="Node Name" name="password">
              <UInput v-model="state.node_name" />
            </UFormGroup>

            <UFormGroup label="Type" name="type">
              <UInput v-model="state.type" />
            </UFormGroup>

            <UFormGroup label="Attribute" name="attribute">
              <UInput v-model="state.attributes" />
            </UFormGroup>
            <UButton type="submit" :disabled="disabled"> Submit </UButton>
          </UForm>
        </UCard>
      </UModal>
    </div>
  </UContainer>
</template>
