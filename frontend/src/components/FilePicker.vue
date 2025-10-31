<script setup lang="ts">
import { ref } from "vue";
import { SelectFile } from "../../wailsjs/go/main/App";
import { ExtractText } from "../../wailsjs/go/main/App";

const props = defineProps<{
  format?: string;
}>();

const emits = defineEmits<{
  (e: "fileLoaded", payload: { path: string; content: string }): void;
}>();

const loading = ref(false);

async function openFile() {
  try {
    loading.value = true;
    // Open system picker; simple call - options vary by Wails version
    const result = (await SelectFile()) as unknown;

    // Normalize result to string[]
    let paths: string[] = [];
    if (Array.isArray(result)) {
      paths = result as string[];
    } else if (typeof result === "string" && result) {
      paths = [result];
    }

    if (paths.length === 0) {
      loading.value = false;
      return;
    }

    const path = paths[0];
    // Call backend to read selected file via generated binding
    const content = (await ExtractText(path)) as string;
    emits("fileLoaded", { path, content });
  } catch (err) {
    console.error("File open/read failed", err);
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <n-button @click="openFile" :loading="loading" type="primary">
    Open File
  </n-button>
</template>
