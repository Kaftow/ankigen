import { defineStore } from "pinia";
import { ref, computed } from "vue";
import * as fileService from "../../services/fileService";

export const useFileManagerStore = defineStore("fileManager", () => {
  const supportedFormats = ref<string[]>([]);
  const lastFile = ref<{ filename: string; content: string } | null>(null);
  const loadingSupportedFormats = ref(false);
  const loadingFile = ref(false);

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
      lastFile.value = res;
      return res;
    } catch (e) {
      console.error("openFile failed", e);
      lastFile.value = null;
      return null;
    } finally {
      loadingFile.value = false;
    }
  }

  function clearFile() {
    lastFile.value = null;
  }

  const hasFile = computed(() => lastFile.value !== null);

  return {
    supportedFormats,
    lastFile,
    loadingSupportedFormats,
    loadingFile,
    loadSupportedFormats,
    openFile,
    clearFile,
    hasFile,
  };
});
