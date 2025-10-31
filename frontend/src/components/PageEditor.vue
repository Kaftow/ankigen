<script setup lang="ts">
import { ref, watch } from "vue";
import VMdEditor from "@kangc/v-md-editor";
import "@kangc/v-md-editor/lib/style/base-editor.css";
import "@kangc/v-md-editor/lib/theme/style/vuepress.css";

const props = defineProps<{
  modelValue: string;
}>();

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
  <n-card>
    <!-- Simple wrapped markdown editor -->
    <VMdEditor v-model="internal" />
  </n-card>
</template>
