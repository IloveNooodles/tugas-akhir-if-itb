<script setup lang="ts">
import type { CompanyUser } from '~/types/company';

interface Props {
  company: CompanyUser;
  pending: boolean;
}

const props = defineProps<Props>();
const { company, pending } = props;
const { name, cluster_name, email, username } = company;
</script>

<template>
  <UCard v-if="pending">
    <h1>Loading</h1>
  </UCard>
  <UCard v-else>
    <h2>
      Welcome,
      <span class="text-bold">
        {{ username }}
      </span>
    </h2>
    <h3 class="text-gray-200">Company: {{ name }}</h3>
    <h3 class="text-gray-400">
      Server:
      <span>
        {{ cluster_name }}
      </span>
    </h3>
    <p class="text-gray-500">Logged in as: {{ email }}</p>
  </UCard>
</template>
