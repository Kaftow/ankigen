import { App } from "vue";
import { createPinia } from "pinia";

const pinia = createPinia();

export function setupStore(app: App) {
  app.use(pinia);
}

export default pinia;

export { useFileManagerStore } from "./modules/fileManager";
