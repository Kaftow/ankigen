<script setup lang="ts">
import { NCard } from "naive-ui";
import ChunkingStrategyConfig from "@/components/ChunkingStrategyConfig.vue";
import ChunkEditor from "@/components/ChunkEditor.vue";
import { useFileStore, useMdBlockStore, useChunkerConfigStore } from "@/stores";

const fileStore = useFileStore();
const mdBlockStore = useMdBlockStore();
const chunkerStore = useChunkerConfigStore();

const handleChunk = async () => {
  if (!chunkerStore.currentStrategy) return;

  if (!fileStore.editedContent) {
    console.error("No content to chunk");
    return;
  }

  await mdBlockStore.fetchChunks(
    fileStore.editedContent,
    chunkerStore.currentStrategy,
    chunkerStore.currentParams,
  );
};
</script>
<template>
  <div class="file-reader">
    <div v-if="!mdBlockStore.hasBlocks" class="controls">
      <ChunkingStrategyConfig @chunk="handleChunk" />
    </div>
    <div v-else class="editor-wrap">
      <n-card title="Chunking Preview">
        <ChunkEditor />
      </n-card>
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
  flex-shrink: 0;
}

.editor-wrap {
  flex: 1;
  min-height: 0;
  overflow: hidden;
}
</style>
