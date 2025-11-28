<script setup lang="ts">
import { ref, computed } from "vue";
import FileEmptyCard from "@/components/FileEmptyCard.vue";
import FilePicker from "../components/FilePicker.vue";
import PageEditor from "../components/PageEditor.vue";
import { useFileManagerStore } from "../stores/modules/fileManager";

const fileStore = useFileManagerStore();

function onFileLoaded(payload: { path?: string; content?: string } | null) {
  if (!payload) {
    fileStore.lastFile = null;
    return;
  }
  fileStore.lastFile = {
    filename: payload.path ?? "",
    content: typeof payload.content === "string" ? payload.content : "",
  };
}

const editorContent = computed<string>({
  get() {
    return fileStore.lastFile?.content ?? "";
  },
  set(value: string) {
    if (fileStore.lastFile) {
      fileStore.lastFile = { ...fileStore.lastFile, content: value };
    } else {
      fileStore.lastFile = { filename: "", content: value };
    }
  },
});
</script>

<template>
  <div class="file-reader">
    <div class="controls">
      <file-picker @fileLoaded="onFileLoaded" />
    </div>

    <FileEmptyCard v-if="!fileStore.lastFile || !fileStore.lastFile.content" />

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
