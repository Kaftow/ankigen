<script setup lang="ts">
import { ref, onMounted, computed, watch } from "vue";
import { GetSupportedExtensions } from "../../wailsjs/go/main/App";

const emits = defineEmits<{
  (e: "update:format", value: string): void;
}>();

const formats = ref<string[]>([]);
const selected = ref<string>("");

onMounted(async () => {
  try {
    const res = (await GetSupportedExtensions()) as string[] | null;
    formats.value = Array.isArray(res) ? res : [];
  } catch (err) {
    console.error("Failed to fetch formats", err);
  }
});

const options = computed(() =>
  formats.value.map((f) => ({ label: f, value: f })),
);

watch(selected, (v) => emits("update:format", v));
</script>

<template>
  <n-select
    v-model:value="selected"
    :options="options"
    placeholder="Select file format"
    style="min-width: 220px"
  />
</template>
