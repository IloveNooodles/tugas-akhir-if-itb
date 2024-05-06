<script setup lang="ts">
import type { CompanyUser } from '~/types/company';

interface Props {
  company: CompanyUser | null;
  pending: boolean;
  error?: any;
}

const props = defineProps<Props>();
</script>

<template>
  <UCard v-if="pending">
    <Loading />
  </UCard>
  <UCard v-else-if="error || !company">
    <div>Sorry, we're having an issue please try again</div>
  </UCard>
  <UCard v-else>
    <h2>
      Welcome,
      <span class="text-bold">
        {{ company.username }}
      </span>
    </h2>
    <h3 class="text-gray-200">Company: {{ company.name }}</h3>
    <h3 class="text-gray-400">
      Server:
      <span>
        {{ company.cluster_name }}
      </span>
    </h3>
    <p class="text-gray-500">Logged in as: {{ company.email }}</p>
  </UCard>
</template>
