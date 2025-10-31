<script setup lang="ts">
import { ref, watch } from "vue";
import PageEditor from "./PageEditor.vue";

const props = defineProps<{
  pages: string[];
}>();

const emits = defineEmits<{
  (e: "update:pages", pages: string[]): void;
}>();

// Local editable copy so child v-model works nicely
const pagesLocal = ref<string[]>([...props.pages]);

watch(
  () => props.pages,
  (v) => {
    pagesLocal.value = [...v];
  },
);

function updatePage(i: number, val: string) {
  pagesLocal.value[i] = val;
  emits("update:pages", [...pagesLocal.value]);
}
</script>

<template>
  <n-tabs :type="'line'">
    <n-tab-pane
      v-for="(p, i) in pagesLocal"
      :key="i"
      :name="String(i)"
      :tab="'Page ' + (i + 1)"
    >
      <page-editor
        v-model="pagesLocal[i]"
        @update:modelValue="(v) => updatePage(i, v)"
      />
    </n-tab-pane>
  </n-tabs>
</template>
