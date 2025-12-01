<script setup lang="ts">
import { computed } from "vue";
import FileEmptyCard from "@/components/FileEmptyCard.vue";
import FilePicker from "@/components/FilePicker.vue";
import PageEditor from "@/components/PageEditor.vue";
import { useFileStore, useWorkflowStore } from "@/stores";

const fileStore = useFileStore();
const workflowStore = useWorkflowStore();

const isEmpty = computed(() => !fileStore.lastFile?.content);

function setStepCompleted(value: boolean) {
  workflowStore.setStepCompleted(value);
}

function onFileLoaded(payload: { filename?: string; content?: string } | null) {
  console.log("File loaded:", payload);
  if (!payload?.content || !payload?.filename) {
    fileStore.lastFile = null;
    return;
  }
  fileStore.lastFile = {
    filename: payload.filename,
    content: payload.content,
  };
  setStepCompleted(true);
}

const editorContent = computed<string>({
  get() {
    return fileStore.lastFile?.content ?? "";
  },
  set(value: string) {
    if (!value) {
      // remove file when empty and notify parent that "next" is disabled
      fileStore.lastFile = null;
      setStepCompleted(false);
      return;
    }
    if (fileStore.lastFile) {
      fileStore.lastFile = { ...fileStore.lastFile, content: value };
    } else {
      fileStore.lastFile = { filename: "", content: value };
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
