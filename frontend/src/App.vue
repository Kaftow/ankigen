<script setup lang="ts">
import { NLayout, NLayoutContent } from "naive-ui";
import StepNavigator from "./components/StepNavigator.vue";
import { useWorkflowStore } from "@/stores";
import { useRouter } from "vue-router";
import { computed } from "vue";

const router = useRouter();
const workflow = useWorkflowStore();

// computed props for StepNavigator
const canNext = computed(() => workflow.isStepCompleted);
const showPrev = computed(() => !workflow.isFirstStep);
const showNext = computed(() => !workflow.isLastStep);

// handle next / prev events
function onNext() {
  workflow.nextStep();
  router.push({ name: workflow.currentStep });
  workflow.setStepCompleted(false); // reset externally
}

function onPrev() {
  workflow.prevStep();
  router.push({ name: workflow.currentStep });
}
</script>

<template>
  <n-layout class="page-container">
    <n-layout-content class="page-content">
      <router-view />
    </n-layout-content>
    <step-navigator
      :canNext="canNext"
      @next="onNext"
      @prev="onPrev"
      :showPrev="showPrev"
      :showNext="showNext"
    />
  </n-layout>
</template>

<style scoped>
.page-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.page-content {
  flex: 1;
  min-height: 0;
  display: flex;
  flex-direction: column;
  padding: 16px;
  overflow: hidden;
}
</style>
