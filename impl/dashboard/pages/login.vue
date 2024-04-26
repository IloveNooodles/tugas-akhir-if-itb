<script setup lang="ts">
import type { FormSubmitEvent } from '#ui/types';
import { login } from '@/api/user/auth';
import {
  type UserLoginSchema as Schema,
  userLoginSchema as schema,
} from '@/types/user';
import { FetchError } from 'ofetch';

const { set } = useAuthStore();
const { $toast } = useNuxtApp();

const state = ref({
  email: '',
  password: '',
});

async function onSubmit(event: FormSubmitEvent<Schema>) {
  const body = event.data;
  try {
    const response = await login(body);
    $toast.success('Success Login, redirecting');
    set(response.data);
    await navigateTo('/');
  } catch (err: any) {
    if (err instanceof FetchError) {
      $toast.error(err.data.message);
    }
  }
}
</script>

<template>
  <UContainer class="flex flex-col justify-center min-h-[100vh] max-w-fit">
    <h1 class="text-center font-sans text-3xl font-bold">
      Remote Deployment Manager
    </h1>
    <UForm
      :schema="schema"
      :state="state"
      class="space-y-4 pt-4"
      @submit="onSubmit"
    >
      <UFormGroup label="Email" name="email">
        <UInput v-model="state.email" />
      </UFormGroup>

      <UFormGroup label="Password" name="password">
        <UInput v-model="state.password" type="password" />
      </UFormGroup>

      <UButton type="submit"> Submit </UButton>
    </UForm>
  </UContainer>
</template>
