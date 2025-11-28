<script setup lang="ts">
import { ref, computed } from "vue";
import { useFileManagerStore } from "@/stores/modules/fileManager";

const emits = defineEmits<{
  (e: "fileLoaded", payload: { filename: string; content: string }): void;
}>();

const fileStore = useFileManagerStore();
const loading = computed(() => fileStore.loadingFile);

async function openFile() {
  try {
    const result = await fileStore.openFile();
    if (result) {
      emits("fileLoaded", result);
    }
  } catch (err) {
    console.error("File open/read failed", err);
  }
}
</script>

<template>
  <n-button @click="openFile" :loading="loading" type="primary">
    Open File
  </n-button>
</template>
