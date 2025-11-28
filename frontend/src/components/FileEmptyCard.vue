<script setup lang="ts">
import { computed, onMounted } from "vue";
import { useFileManagerStore } from "@/stores/modules/fileManager";

const fileStore = useFileManagerStore();

onMounted(() => {
  fileStore.loadSupportedFormats();
});

const options = computed(() =>
  fileStore.supportedFormats.map((fmt) => ({
    label: fmt,
    value: fmt,
  })),
);
</script>

<template>
  <n-card class="file-empty-card">
    <div style="display: flex; flex-direction: column; gap: 12px">
      <div>Open a supported file to view and edit Markdown content.</div>

      <div v-if="options.length">
        <div style="font-size: 12px; color: #666; margin-bottom: 6px">
          Supported formats:
        </div>
        <div style="display: flex; gap: 8px; flex-wrap: wrap">
          <button
            v-for="opt in options"
            :key="opt.value"
            type="button"
            class="format-item"
            style="
              padding: 6px 10px;
              border: 1px solid #d9d9d9;
              border-radius: 6px;
              background: #fff;
              cursor: pointer;
            "
          >
            {{ opt.label }}
          </button>
        </div>
      </div>
    </div>
  </n-card>
</template>

<style scoped>
.format-item {
  padding: 6px 10px;
  border: 1px solid #d9d9d9;
  border-radius: 6px;
  background: #fff;
  cursor: pointer;
}
.file-empty-card {
  margin-top: 20px;
}
</style>
