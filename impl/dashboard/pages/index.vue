<script setup lang="ts">
definePageMeta({
  middleware: 'auth',
});
</script>

<template>
  <UContainer class="flex-1">
    <h1 class="text-center">Dashboard</h1>
    <UDivider />
    <div class="mt-4">
      <h2>Quick Actions</h2>
      <div class="flex gap-2">
        <UButton
          label="Remote Deployment"
          icon="i-heroicons-document-plus"
          class="mt-2 max-w-fit self-end"
          to="/remote-deployment"
          size="xl"
        />
        <UButton
          label="See Devices"
          icon="i-heroicons-document-plus"
          class="mt-2 max-w-fit self-end"
          to="/devices"
          size="xl"
        />
        <UButton
          label="See Groups"
          icon="i-heroicons-document-plus"
          class="mt-2 max-w-fit self-end"
          to="/groups"
          size="xl"
        />
        <UButton
          label="How to register devices"
          icon="i-heroicons-document-plus"
          class="mt-2 max-w-fit self-end"
          to="/faq/how-to-add-devices"
          size="xl"
        />
      </div>
    </div>
  </UContainer>
</template>
