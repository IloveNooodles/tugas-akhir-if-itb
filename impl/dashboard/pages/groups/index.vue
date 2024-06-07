<script setup lang="ts">
import type { FormSubmitEvent } from '#ui/types';
import {
  createGroupSchema as schema,
  type CreateGroupSchema as Schema,
} from '@/types/group';
import { FetchError } from 'ofetch';
import { createGroup, getGroupList } from '~/api/group';

const nuxtApp = useNuxtApp();
const toast = useToast();

const {
  data: groupData,
  error: groupError,
  pending: groupPending,
  refresh: groupRefresh,
} = await getGroupList(nuxtApp);

const columns = computed(() => {
  return generateColumnsFromArray(groupData.value, ['company_id']);
});

const isOpen = ref(false);
const disabled = ref(false);
const state = ref({
  name: '',
});

async function onSubmit(event: FormSubmitEvent<Schema>) {
  const body = event.data;
  try {
    await createGroup(body, nuxtApp);
    toast.add({
      title: 'Success creating groups',
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
    state.value.name = '';
  }
}

async function onDelete() {
  await groupRefresh();
}
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
            <UFormGroup label="Name" name="name">
              <UInput v-model="state.name" type="text" />
            </UFormGroup>
            <UButton type="submit" :disabled="disabled"> Submit </UButton>
          </UForm>
        </UCard>
      </UModal>
    </div>
  </UContainer>
</template>
