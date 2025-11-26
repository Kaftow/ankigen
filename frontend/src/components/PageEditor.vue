<script setup lang="ts">
import { ref, watch, onMounted } from "vue";

import "@kangc/v-md-editor/lib/style/base-editor.css";
import "@kangc/v-md-editor/lib/theme/style/vuepress.css";

import * as Codemirror from "codemirror";
import "codemirror/lib/codemirror.css";
import "codemirror/mode/markdown/markdown";
import "codemirror/mode/javascript/javascript";
import "codemirror/mode/css/css";
import "codemirror/mode/htmlmixed/htmlmixed";
import "codemirror/mode/vue/vue";
import "codemirror/addon/edit/closebrackets";
import "codemirror/addon/edit/closetag";
import "codemirror/addon/edit/matchbrackets";
import "codemirror/addon/display/placeholder";
import "codemirror/addon/selection/active-line";
import "codemirror/addon/scroll/simplescrollbars";
import "codemirror/addon/scroll/simplescrollbars.css";

import VMdEditor from "@kangc/v-md-editor/lib/codemirror-editor";
import githubTheme from "@kangc/v-md-editor/lib/theme/github.js";
import hljs from "highlight.js";

const props = withDefaults(defineProps<{ modelValue?: string }>(), {
  modelValue: "",
});

const emit = defineEmits<{
  (e: "update:modelValue", value: string): void;
}>();

const internal = ref(props.modelValue);
watch(
  () => props.modelValue,
  (v) => {
    if (v !== internal.value) internal.value = v;
  },
);
watch(internal, (v) => emit("update:modelValue", v));

const editorReady = ref(false);

onMounted(() => {
  VMdEditor.Codemirror = (Codemirror as any).default || Codemirror;
  (VMdEditor as any).use(githubTheme, { Hljs: hljs });

  editorReady.value = true;
});
</script>

<template>
  <n-card>
    <div v-if="editorReady">
      <VMdEditor v-model="internal" />
    </div>
  </n-card>
</template>
