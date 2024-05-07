<script setup lang="ts">
import type { User } from '~/types/user';

interface Props {
  user: User;
  pending: boolean;
}

const props = defineProps<Props>();
const { user, pending } = props;
const { name, created_at, email, updated_at } = user;

const createdAt = formatDate(created_at);
const updatedAt = formatDate(updated_at);
</script>

<template>
  <UCard v-if="pending">
    <h1>Loading</h1>
  </UCard>
  <UCard v-else>
    <h2 class="font-bold">Welcome {{ name }}</h2>
    <p class="">{{ email }}</p>
    <p>Created At: {{ createdAt }}</p>
    <p>Updated At: {{ updatedAt }}</p>
  </UCard>
</template>
