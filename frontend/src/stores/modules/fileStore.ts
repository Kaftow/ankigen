import { defineStore } from "pinia";
import { ref, computed, watch } from "vue";
import * as fileService from "@/services/fileService";

export const useFileStore = defineStore("fileStore", () => {
  const supportedFormats = ref<string[]>([]);
  const filename = ref<string | null>(null);
  const rawContent = ref<string | null>(null);
  const editedContent = ref<string | null>(null);
  const loadingSupportedFormats = ref(false);
  const loadingFile = ref(false);

  watch(rawContent, (newVal) => {
    editedContent.value = newVal;
  });

  async function loadSupportedFormats() {
    loadingSupportedFormats.value = true;
    try {
      supportedFormats.value = await fileService.getSupportedExtensions();
    } catch (e) {
      console.error("loadFormats failed", e);
      supportedFormats.value = [];
    } finally {
      loadingSupportedFormats.value = false;
    }
  }

  async function openFile() {
    loadingFile.value = true;
    try {
      const res = await fileService.pickAndRead();
      if (res) {
        filename.value = res.filename;
        rawContent.value = res.content;
        return res;
      }
      filename.value = null;
      rawContent.value = null;
      return null;
    } catch (e) {
      console.error("openFile failed", e);
      filename.value = null;
      rawContent.value = null;
      return null;
    } finally {
      loadingFile.value = false;
    }
  }

  function clearFile() {
    filename.value = null;
    rawContent.value = null;
  }

  const hasFile = computed(
    () => filename.value !== null && rawContent.value !== null,
  );

  return {
    supportedFormats,
    filename,
    rawContent,
    editedContent,
    loadingSupportedFormats,
    loadingFile,
    loadSupportedFormats,
    openFile,
    clearFile,
    hasFile,
  };
});
