<script setup lang="ts">
import { getGroupList } from '~/api/group';

const nuxtApp = useNuxtApp();

const {
  data: groupData,
  error: groupError,
  pending: groupPending,
} = await getGroupList(nuxtApp);

const columns = computed(() => {
  return generateColumnsFromArray(groupData.value, [
    // 'created_at', 'updated_at'
  ]);
});
</script>

<template>
  <UContainer>
    <h1 class="text-center">Groups</h1>
    <UDivider />
    <div class="wrap">
      <h2>Available Groups</h2>
      <GroupList
        :data="groupData"
        :pending="groupPending"
        :columns="columns"
        :error="groupError"
      />
      <UButton
        label="Add Group"
        icon="i-heroicons-plus-solid"
        class="mt-2 max-w-fit self-end"
      />
    </div>
  </UContainer>
</template>
