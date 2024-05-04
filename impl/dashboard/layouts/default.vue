<script setup lang="ts"></script>

<template>
  <main class="flex">
    <Sidebar />

    <div class="content flex flex-col flex-1 bg-slate-950 justify-between">
      <slot class="w-full" />
      <Footer />
    </div>
  </main>
</template>
