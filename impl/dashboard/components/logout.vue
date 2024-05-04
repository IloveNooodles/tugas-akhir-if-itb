<script setup lang="ts">
async function logout() {
  await navigateTo('/login');
}
</script>

<template>
  <UButton color="gray" label="Logout" variant="ghost" @click="logout"/>
</template>
