<script setup lang="ts">
import { useStorage } from '@vueuse/core';

async function logout() {
  useStorage('accessToken', '').value = '';
  useStorage('refreshToken', '').value = '';
  await navigateTo('/login');
}
</script>

<template>
  <UButton color="gray" label="Logout" variant="ghost" @click="logout" />
</template>
