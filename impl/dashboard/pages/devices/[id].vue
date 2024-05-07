<script setup lang="ts">
import { getDeviceById, getDeviceGroupsById } from '~/api/device';

const route = useRoute();
const nuxtApp = useNuxtApp();
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

const columns = computed(() => {
  return generateColumnsFromArray(groupData.value, ['company_id']);
});
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
          <UButton icon="i-heroicons-pencil-square-20-solid" />
          <UButton icon="i-heroicons-trash-20-solid" />
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
        <div>
          <h2>Groups</h2>
          <GroupList
            :data="groupData"
            :pending="groupPending"
            :columns="columns"
            :error="groupError"
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
