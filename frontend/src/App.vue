<script setup lang="ts">
import { NDialogProvider, NMessageProvider } from "naive-ui";
import StepNavigator from "@/components/StepNavigator.vue";
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
  <n-message-provider>
    <n-dialog-provider>
      <div class="page-root">
        <div class="page-header">
          <h1 class="page-title">{{ workflow.pageTitle }}</h1>
        </div>

        <n-scrollbar class="page-scroll">
          <router-view />
        </n-scrollbar>

        <div class="page-footer">
          <step-navigator
            :canNext="canNext"
            :showPrev="showPrev"
            :showNext="showNext"
            @next="onNext"
            @prev="onPrev"
          />
        </div>
      </div>
    </n-dialog-provider>
  </n-message-provider>
</template>

<style scoped>
@font-face {
  font-family: "Montserrat";
  src: url("@/assets/fonts/Montserrat-SemiBold.woff2") format("woff2");
  font-weight: 600;
  font-display: swap;
}

.page-root {
  height: 100vh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  background-color: #fff;
}

.page-header {
  background-color: #fff;
  padding: 24px 32px;
  border-bottom: 1px solid #e5e7eb;
  display: flex;
  align-items: center;
}

.page-title {
  margin: 0;
  font-size: 28px;
  font-weight: 600;
  color: #1f2937;
  font-family: "Montserrat", sans-serif;
}

.page-scroll {
  flex: 1;
  width: 100%;
  background-color: #fff;
  min-height: 0;
}

.page-scroll :deep(.n-scrollbar) {
  background-color: #fff;
}

.page-scroll :deep(.n-scrollbar-content) {
  background-color: #fff;
}

.page-footer {
  flex-shrink: 0;
  width: 100%;
  min-height: 60px;
  padding: 16px;
  border-top: 1px solid #e0e0e0;
  background-color: #fff;
  box-sizing: border-box;
  display: flex;
  align-items: center;
}
</style>
