<script setup lang="ts">
import { ref, watch } from "vue";
import { MdEditor, MdPreview } from "md-editor-v3";
import "md-editor-v3/lib/style.css";

const props = withDefaults(
  defineProps<{ modelValue?: string; previewMode?: boolean }>(),
  { modelValue: "", previewMode: false },
);

const emit = defineEmits<{
  (e: "update:modelValue", value: string): void;
}>();

const internal = ref(props.modelValue);
watch(
  () => props.modelValue,
  (v) => {
    if (v !== internal.value) internal.value = v;
  },
);
watch(internal, (v) => emit("update:modelValue", v));
</script>

<template>
  <n-card class="editor-card">
    <div class="editor-wrapper">
      <n-scrollbar
        trigger="always"
        :scrollbar-style="{ right: '0px', width: '8px', opacity: 1 }"
      >
        <!-- Toggle between editor and preview -->
        <template v-if="!props.previewMode">
          <MdEditor v-model="internal" class="mdeditor" />
        </template>
        <template v-else>
          <MdPreview :modelValue="internal" class="mdeditor" />
        </template>
      </n-scrollbar>
    </div>
  </n-card>
</template>

<style scoped>
.editor-card {
  height: 100%;
  display: flex;
  flex-direction: column;
  min-height: 0;
}

.editor-wrapper {
  width: 100%;
  max-height: 600px;
  overflow: hidden;
}

::v-deep(.n-scrollbar) {
  width: 100%;
  height: 100%;
  overflow-y: auto;
}

.mdeditor {
  width: 100%;
  min-height: 400px;
}

::v-deep(.md-editor) {
  width: 100% !important;
  overflow: visible !important;
  min-height: 400px !important;
}
</style>
