<script setup lang="ts">
definePageMeta({
  middleware: 'auth',
});
</script>

<template>
  <UContainer class="flex-1">
    <h1 class="text-center">Dashboard</h1>
    <UDivider />
  </UContainer>
</template>
