<script setup lang="ts"></script>

<template>
  <UContainer>
    <h1 class="text-center">FAQ</h1>
    <UDivider />
    <p class="text-left pt-4">
      These are most common question when using this service
    </p>
    <ul class="p-4 pt-0">
      <li>
        <ULink
          to="/faq/how-to-add-devices"
          inactive-class="text-gray-500 dark:text-gray-400 hover:text-gray-700 dark:hover:text-gray-200"
        >
          How to add devices
        </ULink>
      </li>
    </ul>
  </UContainer>
</template>

<style scoped>
li {
  list-style: circle;
}
</style>
