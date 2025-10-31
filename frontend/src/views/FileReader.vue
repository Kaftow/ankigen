<script setup lang="ts">
import { ref } from "vue";
import FormatSelector from "../components/FormatSelector.vue";
import FilePicker from "../components/FilePicker.vue";
import PageTabs from "../components/PageTabs.vue";

const selectedFormat = ref<string | null>(null);
const rawContent = ref<string>("");
const pages = ref<string[]>([]);

function onFormatUpdate(format: string) {
  selectedFormat.value = format || null;
}

function onFileLoaded(payload: { path: string; content: string }) {
  rawContent.value = payload.content || "";
  pages.value = splitIntoPages(rawContent.value);
}

// Simple splitter: group paragraphs until approx maxCharsPerPage
function splitIntoPages(text: string, maxCharsPerPage = 3000) {
  if (!text) return [];
  const paragraphs = text
    .split(/\n{2,}/)
    .map((p) => p.trim())
    .filter(Boolean);
  const out: string[] = [];
  let cur = "";

  for (const p of paragraphs) {
    if ((cur + "\n\n" + p).length > maxCharsPerPage) {
      if (cur.trim()) out.push(cur.trim());
      cur = p;
    } else {
      cur = cur ? cur + "\n\n" + p : p;
    }
  }
  if (cur.trim()) out.push(cur.trim());
  // If no paragraphs found but text exists fallback
  if (out.length === 0 && text.trim()) out.push(text.trim());
  return out;
}

function updatePages(newPages: string[]) {
  pages.value = newPages;
}
</script>

<template>
  <div style="padding: 16px; display: grid; gap: 12px">
    <div style="display: flex; gap: 12px; align-items: center">
      <format-selector @update:format="onFormatUpdate" />
      <file-picker
        :format="selectedFormat || undefined"
        @fileLoaded="onFileLoaded"
      />
    </div>

    <div v-if="pages.length === 0" style="margin-top: 20px">
      <n-card
        >Open a supported file to view and edit paginated Markdown
        content.</n-card
      >
    </div>

    <div v-else>
      <page-tabs :pages="pages" @update:pages="updatePages" />
    </div>
  </div>
</template>
