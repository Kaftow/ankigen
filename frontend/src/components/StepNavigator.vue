<script setup lang="ts">
import { NButton, NSpace } from "naive-ui";

const props = defineProps<{
  canNext: boolean;
  showPrev: boolean;
  showNext: boolean;
}>();

// emits next / prev events
const emit = defineEmits<{
  (e: "next"): void;
  (e: "prev"): void;
}>();

// actions
function handlePrev() {
  emit("prev");
}

function handleNext() {
  if (!props.canNext) return;
  emit("next");
}
</script>

<template>
  <n-space justify="end">
    <n-button v-if="props.showPrev" @click="handlePrev" strong>
      Previous
    </n-button>

    <n-button
      v-if="props.showNext"
      type="primary"
      :disabled="!props.canNext"
      @click="handleNext"
      strong
    >
      Next
    </n-button>
  </n-space>
</template>
