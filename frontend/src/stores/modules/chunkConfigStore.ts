import { defineStore } from "pinia";
import { ref, computed } from "vue";
import { z } from "zod";

const ParamSchema = z.object({
  name: z.string(),
  type: z.enum(["integer", "string", "boolean"]),
  default: z.any(),
  description: z.string(),
  min: z.number().optional(),
  max: z.number().optional(),
  required: z.boolean().optional(),
  enum: z.array(z.string()).optional(),
});

const StrategyDefinitionZod = z.object({
  description: z.string(),
  params: z.array(ParamSchema),
});

type StrategyDefinition = z.infer<typeof StrategyDefinitionZod>;

export const useChunkerConfigStore = defineStore("chunker", () => {
  const currentStrategy = ref<string | null>(null);
  const currentParams = ref<Record<string, any>>({});

  // Strategy definitions dictionary
  const strategyDefinitions: Record<string, StrategyDefinition> = {
    fixedLength: {
      description: "Fixed length splitting strategy",
      params: [
        {
          name: "maxChars",
          type: "integer",
          default: 512,
          min: 1,
          description: "Maximum characters per chunk",
        },
      ],
    },
    token: {
      description: "Token-based splitting strategy",
      params: [
        {
          name: "maxTokens",
          type: "integer",
          default: 512,
          min: 1,
          description: "Maximum tokens per chunk",
        },
        {
          name: "encodingName",
          type: "string",
          default: "cl100k_base",
          description: "Encoding name for tokenizer",
          enum: ["o200k_base", "cl100k_base", "p50k_base", "r50k_base"],
        },
      ],
    },
  };

  const updateCurrentParams = (updates: Partial<Record<string, any>>) => {
    currentParams.value = { ...currentParams.value, ...updates };
  };

  // Get all available strategies
  const getAvailableStrategies = computed(() => {
    return Object.entries(strategyDefinitions).map(([name, definition]) => ({
      name,
      description: definition.description,
    }));
  });

  // Get current strategy parameter schema definitions
  const getStrategyParamSchema = computed(() => {
    if (!currentStrategy.value) return [];
    return strategyDefinitions[currentStrategy.value]?.params || [];
  });

  // Validate strategy params
  const validateParams = (
    strategyName: string,
    params: Record<string, any>,
  ): boolean => {
    const paramSchemas = strategyDefinitions[strategyName]?.params || [];
    for (const schema of paramSchemas) {
      try {
        if (schema.type === "integer") {
          const value = params[schema.name];
          if (value === undefined && !schema.required) continue;
          if (typeof value !== "number" || !Number.isInteger(value))
            return false;
          if (schema.min !== undefined && value < schema.min) return false;
          if (schema.max !== undefined && value > schema.max) return false;
        } else if (schema.type === "string") {
          const value = params[schema.name];
          if (value === undefined && !schema.required) continue;
          if (typeof value !== "string") return false;
          if (schema.enum && !schema.enum.includes(value)) return false;
        }
      } catch {
        return false;
      }
    }
    return true;
  };

  return {
    currentStrategy,
    currentParams,
    updateCurrentParams,
    getAvailableStrategies,
    getStrategyParamSchema,
    validateParams,
  };
});
