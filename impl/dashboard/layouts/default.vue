<script setup lang="ts"></script>

<template>
  <main class="flex">
    <Sidebar />

    <div
      class="content flex flex-col flex-1 bg-slate-950 justify-between max-w-full width"
    >
      <slot />
      <Footer />
    </div>
  </main>
</template>

<style scoped>
.width {
  max-width: calc(100% - 165px);
}
</style>
