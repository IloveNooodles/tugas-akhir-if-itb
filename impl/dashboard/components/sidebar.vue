<script setup lang="ts">
const links = [
  {
    label: 'Dashboard',
    icon: 'i-heroicons-home',
    to: '/',
  },
  {
    label: 'Account',
    icon: 'i-heroicons-user-group',
    to: '/account',
  },
  {
    label: 'Devices',
    icon: 'i-heroicons-device-phone-mobile',
    to: '/devices',
  },
  {
    label: 'Groups',
    icon: 'i-heroicons-archive-box',
    to: '/groups',
  },
  {
    label: 'Deployments',
    icon: 'i-heroicons-command-line',
    to: '/deployments',
  },
  {
    label: 'FAQ',
    icon: 'i-heroicons-question-mark-circle',
    to: '/faq',
  },
];
</script>

<template>
  <section class="flex flex-col justify-between min-h-[100vh] p-4 bg-[#0e172a]">
    <UVerticalNavigation :links="links" />
    <Logout />
  </section>
</template>

<style></style>
