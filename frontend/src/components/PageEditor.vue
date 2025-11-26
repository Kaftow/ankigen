<script setup lang="ts">
import { ref, watch } from "vue";
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";

const props = withDefaults(defineProps<{ modelValue?: string }>(), {
  modelValue: "",
});

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
    <div class="editor-inner">
      <MdEditor v-model="internal" class="mdeditor-fill" />
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

.editor-inner {
  flex: 1;
  min-height: 0;
  overflow: hidden;
}

.mdeditor-fill {
  height: 100%;
}

::v-deep .md-editor {
  height: 100% !important;
  min-height: 0 !important;
}
</style>
