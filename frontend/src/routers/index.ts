import { createRouter, createWebHistory } from "vue-router";
import FileLoader from "@/views/FileLoader.vue";
import TextEditor from "@/views/TextEditor.vue";
import TextChunker from "@/views/TextChunker.vue";

export const stepRouteMap: Record<string, string> = {
  fileLoader: "/file-loader",
  textEditor: "/text-editor",
  textChunker: "/text-chunker",
};

const routes = [
  { path: "/", redirect: "/file-loader" },
  { path: "/file-loader", name: "fileLoader", component: FileLoader },
  { path: "/text-editor", name: "textEditor", component: TextEditor },
  { path: "/text-chunker", name: "textChunker", component: TextChunker },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
