<script setup lang="ts">
import { getDeviceList } from '~/api/device';

const nuxtApp = useNuxtApp();

const {
  data: devicesData,
  error: devicesError,
  pending: devicesPending,
} = await getDeviceList(nuxtApp);

const columns = computed(() => {
  return generateColumnsFromArray(devicesData.value, [
    // 'created_at', 'updated_at'
  ]);
});
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
      />
    </div>
  </UContainer>
</template>
