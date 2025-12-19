<script setup lang="ts">
import { computed } from "vue";
import {
  NCard,
  NSelect,
  NFormItem,
  NForm,
  NInputNumber,
  NInput,
  NCheckbox,
  NButton,
  NSpace,
} from "naive-ui";
import { useChunkerConfigStore } from "@/stores";

const chunkerStore = useChunkerConfigStore();

const emit = defineEmits<{
  (e: "chunk"): void;
}>();

const strategyOptions = computed(() => {
  try {
    const strategies = chunkerStore.getAvailableStrategies;
    return Array.isArray(strategies)
      ? strategies.map((strategy) => ({
          label: strategy.description,
          value: strategy.name,
        }))
      : [];
  } catch (e) {
    console.error("Error getting strategies:", e);
    return [];
  }
});

const onStrategyChange = () => {
  // Reset parameters and set default values
  try {
    chunkerStore.currentParams = {};
    const schema = chunkerStore.getStrategyParamSchema;
    if (Array.isArray(schema)) {
      schema.forEach((param) => {
        chunkerStore.currentParams[param.name] = param.default;
      });
    }
  } catch (e) {
    console.error("Error changing strategy:", e);
  }
};

const handleChunk = () => {
  try {
    if (!chunkerStore.currentStrategy) {
      console.error("No strategy selected, cannot chunk");
      return;
    }

    const isValid_ = chunkerStore.validateParams(
      chunkerStore.currentStrategy,
      chunkerStore.currentParams,
    );

    if (!isValid_) {
      console.error("Invalid parameters, cannot chunk");
      return;
    }

    emit("chunk");
    console.log("Chunk executed");
  } catch (e) {
    console.error("Chunk error:", e);
  }
};
</script>

<template>
  <div style="width: 100%">
    <n-card title="Chunker Configuration">
      <n-form :model="{ strategy: chunkerStore.currentStrategy }">
        <n-form-item
          label="Chunking Strategy"
          path="strategy"
          label-placement="left"
        >
          <n-select
            v-model:value="chunkerStore.currentStrategy"
            :options="strategyOptions"
            @update:value="onStrategyChange"
            placeholder="Select a strategy"
            style="width: 100%"
          />
        </n-form-item>
      </n-form>

      <!-- Parameters -->
      <n-form
        v-if="chunkerStore.currentStrategy"
        :model="chunkerStore.currentParams"
      >
        <n-form-item
          v-for="param in chunkerStore.getStrategyParamSchema"
          :key="param.name"
          :label="param.description"
          :path="param.name"
          label-placement="left"
        >
          <!-- Integer input -->
          <n-input-number
            v-if="param.type === 'integer'"
            v-model:value="chunkerStore.currentParams[param.name]"
            :min="param.min"
            :max="param.max"
            :step="1"
            :default-value="param.default"
            :placeholder="`Default: ${param.default}`"
          />

          <!-- Enum select -->
          <n-select
            v-else-if="param.type === 'string' && param.enum"
            v-model:value="chunkerStore.currentParams[param.name]"
            :options="param.enum.map((e) => ({ label: e, value: e }))"
            :default-value="param.default"
          />

          <!-- String input -->
          <n-input
            v-else-if="param.type === 'string'"
            v-model:value="chunkerStore.currentParams[param.name]"
            type="text"
            :placeholder="`Default: ${param.default}`"
          />

          <!-- Boolean checkbox -->
          <n-checkbox
            v-else-if="param.type === 'boolean'"
            v-model:checked="chunkerStore.currentParams[param.name]"
            :default-checked="param.default"
          >
            {{ param.description }}
          </n-checkbox>
        </n-form-item>
      </n-form>

      <!-- Action buttons -->
      <n-space v-if="chunkerStore.currentStrategy" style="margin-top: 1rem">
        <n-button type="primary" @click="handleChunk"> Execute Chunk </n-button>
      </n-space>
    </n-card>
  </div>
</template>
