<script setup lang="ts" generic="T">
import { FetchError } from 'ofetch';
import { deleteGroupDeviceByID } from '~/api/groupdevice';

interface Props {
  data: T[] | null;
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
const route = useRoute();
const nuxtApp = useNuxtApp();
const toast = useToast();
const id = route.params.id as string;

const props = defineProps<Props>();
const emits = defineEmits(['onDelete']);
const disabled = ref(false);

async function deleteByID(id: string) {
  try {
    disabled.value = true;
    await deleteGroupDeviceByID(id, nuxtApp);
    emits('onDelete');
    toast.add({ title: `Success deleting ${id}` });
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
        label: 'Delete',
        icon: 'i-heroicons-trash-20-solid',
        disabled: disabled.value,
        click: async () => {
          await deleteByID(row.id);
        },
      },
    ],
  ];
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
