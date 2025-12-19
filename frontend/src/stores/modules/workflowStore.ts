import { defineStore } from "pinia";
import { ref, computed } from "vue";

// step key type
export type StepKey = string;

export const useWorkflowStore = defineStore("workflow", () => {
  // ordered steps
  const steps = ref<StepKey[]>(["fileLoader", "textEditor", "textChunker"]);

  // current step index
  const currentIndex = ref(0);

  // current step key
  const currentStep = computed(() => steps.value[currentIndex.value]);

  // computed state helpers
  const isFirstStep = computed(() => currentIndex.value === 0);
  const isLastStep = computed(
    () => currentIndex.value === steps.value.length - 1,
  );

  const isStepCompleted = ref(false);

  function setStepCompleted(v: boolean) {
    isStepCompleted.value = v;
  }

  // navigation
  function nextStep() {
    if (!isLastStep.value) currentIndex.value++;
  }

  function prevStep() {
    if (!isFirstStep.value) currentIndex.value--;
  }

  function goToStep(step: StepKey) {
    const index = steps.value.indexOf(step);
    if (index !== -1) currentIndex.value = index;
  }

  function reset() {
    currentIndex.value = 0;
  }

  return {
    steps,
    currentIndex,
    currentStep,
    nextStep,
    prevStep,
    goToStep,
    isFirstStep,
    isLastStep,
    isStepCompleted,
    setStepCompleted,
    reset,
  };
});
