<script setup lang="ts">
import { computed, watchEffect } from "vue";
import PageEditor from "@/components/PageEditor.vue";
import { useFileStore, useWorkflowStore } from "@/stores";

const fileStore = useFileStore();
const workflowStore = useWorkflowStore();

// Mark step completed when edited content is non-empty
watchEffect(() => {
  const content = fileStore.editedContent;
  const stepCompleted = workflowStore.isStepCompleted;

  // Only mark step completed if not already completed and content is non-empty
  if (!stepCompleted && content && content.trim() !== "") {
    workflowStore.setStepCompleted(true);
  }
});

const editorContent = computed<string>({
  get() {
    return fileStore.editedContent ?? "";
  },
  set(value: string) {
    if (!value) {
      // Remove file when empty and notify parent that "next" is disabled
      fileStore.clearFile();
      return;
    }
    if (fileStore.editedContent) {
      fileStore.editedContent = value;
    } else {
      fileStore.editedContent = value;
    }
  },
});
</script>

<template>
  <div class="file-reader">
    <div class="editor-wrap">
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
