<script setup lang="ts">
import { computed, watchEffect } from "vue";
import FileEmptyCard from "@/components/FileEmptyCard.vue";
import FilePicker from "@/components/FilePicker.vue";
import PageEditor from "@/components/PageEditor.vue";
import { useFileStore, useWorkflowStore } from "@/stores";

const fileStore = useFileStore();
const workflowStore = useWorkflowStore();

const isEmpty = computed(() => !fileStore.hasFile);

// Mark step completed when raw content is non-empty
watchEffect(() => {
  const content = fileStore.rawContent;
  const stepCompleted = workflowStore.isStepCompleted;

  // Only mark step completed if not already completed and content is non-empty
  if (!stepCompleted && content && content.trim() !== "") {
    workflowStore.setStepCompleted(true);
  }
});

function onFileLoaded(payload: { filename?: string; content?: string } | null) {
  console.log("File loaded:", payload);
  if (!payload?.content || !payload?.filename) {
    fileStore.clearFile();
    return;
  }
  fileStore.filename = payload.filename;
  fileStore.rawContent = payload.content;
}

const readOnlyContent = computed<string>(() => fileStore.rawContent ?? "");
</script>

<template>
  <div class="file-reader">
    <div class="controls">
      <file-picker @fileLoaded="onFileLoaded" />
    </div>

    <FileEmptyCard v-if="isEmpty" />

    <div v-else class="editor-wrap">
      <page-editor v-model="readOnlyContent" :previewMode="true" />
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
