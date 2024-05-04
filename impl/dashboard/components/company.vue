<script setup lang="ts">
import type { Company } from '~/types/company';

interface Props {
  company: Company;
}

defineProps<Props>();
</script>

<template>
  <UCard>
    <h2>{{ company.name }}</h2>
  </UCard>
</template>
