<script setup lang="ts">
import { NCard } from "naive-ui";
import { computed, watchEffect } from "vue";
import ChunkingStrategyConfig from "@/components/ChunkingStrategyConfig.vue";
import BlockEditor from "@/components/BlockEditor.vue";
import {
  useFileStore,
  useMdBlockStore,
  useChunkerConfigStore,
  useWorkflowStore,
} from "@/stores";

const fileStore = useFileStore();
const mdBlockStore = useMdBlockStore();
const chunkerStore = useChunkerConfigStore();
const workflowStore = useWorkflowStore();

const isBlockEmpty = computed(() => !mdBlockStore.hasBlocks);

// Mark step completed when raw content is non-empty
watchEffect(() => {
  const stepCompleted = workflowStore.isStepCompleted;

  // Only mark step completed if not already completed and content is non-empty
  if (!stepCompleted && !isBlockEmpty.value) {
    workflowStore.setStepCompleted(true);
  }
  // Unmark step completed if blocks become empty
  if (stepCompleted && isBlockEmpty.value) {
    workflowStore.setStepCompleted(false);
  }
});

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
        <BlockEditor />
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
