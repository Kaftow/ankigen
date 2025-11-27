<script setup lang="ts">
import FileReader from "./views/FileReader.vue";
</script>

<template>
  <n-layout class="page-container">
    <n-layout-content class="page-content">
      <file-reader />
    </n-layout-content>
  </n-layout>
</template>

<style scoped>
.page-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.page-content {
  flex: 1;
  min-height: 0;
  display: flex;
  flex-direction: column;
  padding: 16px;
  overflow: hidden;
}
</style>
