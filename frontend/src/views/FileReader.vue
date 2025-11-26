<script setup lang="ts">
import { ref } from "vue";
import FormatSelector from "../components/FormatSelector.vue";
import FilePicker from "../components/FilePicker.vue";
import PageEditor from "../components/PageEditor.vue";

const selectedFormat = ref<string | null>(null);
const rawContent = ref<string>("");

function onFormatUpdate(format: string) {
  selectedFormat.value = format || null;
}

function onFileLoaded(payload: { path: string; content: string }) {
  rawContent.value =
    typeof payload?.content === "string" ? payload.content : "";
}
</script>

<template>
  <div style="padding: 16px; display: grid; gap: 12px">
    <div style="display: flex; gap: 12px; align-items: center">
      <format-selector @update:format="onFormatUpdate" />
      <file-picker
        :format="selectedFormat || undefined"
        @fileLoaded="onFileLoaded"
      />
    </div>

    <div v-if="!rawContent" style="margin-top: 20px">
      <n-card>Open a supported file to view and edit Markdown content.</n-card>
    </div>

    <div v-else>
      <page-editor v-model="rawContent" />
    </div>
  </div>
</template>
