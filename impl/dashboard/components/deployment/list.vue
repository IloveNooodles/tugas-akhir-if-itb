<script setup lang="ts">
import { FetchError } from 'ofetch';
import { deleteDeploymentByID } from '~/api/deployment';
import type { Deployment } from '~/types/deployment';

interface Props {
  data: Deployment[] | null;
  pending: boolean;
  columns: {
    [key: string]: any;
    key: string;
    sortable?: boolean | undefined;
    sort?: ((a: any, b: any, direction: 'asc' | 'desc') => number) | undefined;
    direction?: 'asc' | 'desc' | undefined;
    class?: string | undefined;
  }[];
  error?: any;
}

const nuxtApp = useNuxtApp();
const toast = useToast();

const props = defineProps<Props>();
const emits = defineEmits(['onDelete']);
const disabled = ref(false);

async function deleteByID(id: string) {
  try {
    disabled.value = true;
    await deleteDeploymentByID(id, nuxtApp);
    emits('onDelete');
    toast.add({ title: `Success deleting deployments ${id}` });
  } catch (err: any) {
    if (err instanceof FetchError && err.data) {
      toast.add({ title: err.data.message, color: 'red' });
      return;
    }

    toast.add({ title: 'Please try again', color: 'red' });
  } finally {
    disabled.value = false;
  }
}

const dropdownItems = computed(() => {
  return (row: any) => [
    [
      {
        label: 'Detail',
        icon: 'i-heroicons-document-magnifying-glass-16-solid',
        click: () => {
          navigateTo(`/deployments/${row.id}`);
        },
      },
    ],
    [
      {
        label: 'Delete',
        icon: 'i-heroicons-trash-20-solid',
        click: async () => {
          await deleteByID(row.id);
        },
      },
    ],
  ];
});

const columns = computed(() => {
  return generateColumnsFromArray(props.data, [
    'created_at',
    'updated_at',
    'company_id',
  ]);
});
</script>

<template>
  <UCard v-if="pending">
    <Loading />
  </UCard>
  <UCard v-else-if="error || !data">
    <div>Sorry, we're having an issue please try again</div>
  </UCard>
  <Table
    v-else
    :data="data"
    :pending="pending"
    :columns="columns"
    :error="error"
    :dropdown-items="dropdownItems"
  >
  </Table>
</template>
