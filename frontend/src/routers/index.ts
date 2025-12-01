import { createRouter, createWebHistory } from "vue-router";
import FileLoader from "@/views/FileLoader.vue";
import TextEditor from "@/views/TextEditor.vue";

export const stepRouteMap: Record<string, string> = {
  fileLoader: "/file-loader",
  textEditor: "/text-editor",
};

const routes = [
  { path: "/", redirect: "/file-loader" },
  { path: "/file-loader", name: "fileLoader", component: FileLoader },
  { path: "/text-editor", name: "textEditor", component: TextEditor },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
