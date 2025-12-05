import { App } from "vue";
import { createPinia } from "pinia";

const pinia = createPinia();

export function setupStore(app: App) {
  app.use(pinia);
}

export default pinia;

export { useFileStore } from "./modules/fileStore";
export { useWorkflowStore } from "./modules/workflowStore";
export { useMdBlockStore } from "./modules/mdBlockStore";
