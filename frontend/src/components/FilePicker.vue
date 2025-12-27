<script setup lang="ts">
import { computed } from "vue";
import { useFileStore } from "@/stores";
import { useNotification } from "@/composables/useNotification";

const emits = defineEmits<{
  (e: "fileLoaded", payload: { filename: string; content: string }): void;
}>();

const fileStore = useFileStore();
const { showError, showInfo } = useNotification();
const loading = computed(() => fileStore.loadingFile);

async function openFile() {
  try {
    showInfo("Opening file dialog...");
    const result = await fileStore.openFile();
    if (result) {
      emits("fileLoaded", result);
    }
  } catch (err) {
    showError(
      `File open/read failed: ${err instanceof Error ? err.message : "Unknown error"}`,
    );
  }
}
</script>

<template>
  <n-button @click="openFile" :loading="loading" type="primary">
    Open File
  </n-button>
</template>
