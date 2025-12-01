<script setup lang="ts">
import { computed } from "vue";
import FileEmptyCard from "@/components/FileEmptyCard.vue";
import FilePicker from "@/components/FilePicker.vue";
import PageEditor from "@/components/PageEditor.vue";
import { useFileStore, useWorkflowStore } from "@/stores";

const fileStore = useFileStore();
const workflowStore = useWorkflowStore();

const isEmpty = computed(() => !fileStore.hasFile);

function setStepCompleted(value: boolean) {
  workflowStore.setStepCompleted(value);
}

function onFileLoaded(payload: { filename?: string; content?: string } | null) {
  console.log("File loaded:", payload);
  if (!payload?.content || !payload?.filename) {
    fileStore.clearFile();
    setStepCompleted(false);
    return;
  }
  fileStore.filename = payload.filename;
  fileStore.rawContent = payload.content;
  setStepCompleted(true);
}

const editorContent = computed<string>({
  get() {
    return fileStore.rawContent ?? "";
  },
  set(value: string) {
    if (!value) {
      // remove file when empty and notify parent that "next" is disabled
      fileStore.clearFile();
      setStepCompleted(false);
      return;
    }
    if (fileStore.rawContent) {
      fileStore.rawContent = value;
    } else {
      fileStore.rawContent = value;
    }
    setStepCompleted(true);
  },
});
</script>

<template>
  <div class="file-reader">
    <div class="controls">
      <file-picker @fileLoaded="onFileLoaded" />
    </div>

    <FileEmptyCard v-if="isEmpty" />

    <div v-else class="editor-wrap">
      <page-editor v-model="editorContent" />
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
