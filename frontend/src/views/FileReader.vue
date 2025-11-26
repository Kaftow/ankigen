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
  <div class="file-reader">
    <div class="controls">
      <format-selector @update:format="onFormatUpdate" />
      <file-picker
        :format="selectedFormat || undefined"
        @fileLoaded="onFileLoaded"
      />
    </div>

    <div v-if="!rawContent" class="empty">
      <n-card>Open a supported file to view and edit Markdown content.</n-card>
    </div>

    <div v-else class="editor-wrap">
      <page-editor v-model="rawContent" />
    </div>
  </div>
</template>

<style scoped>
.file-reader {
  height: 100%;
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  min-height: 0;
}

.controls {
  display: flex;
  gap: 12px;
  align-items: center;
}

.empty {
  margin-top: 20px;
}

.editor-wrap {
  flex: 1;
  min-height: 0;
  overflow: hidden;
}
</style>
