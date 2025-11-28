import { createApp } from "vue";
import App from "./App.vue";
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import naive from "naive-ui";
import { setupStore } from "@/stores/index";

const app = createApp(App);
setupStore(app); 
app.use(naive);
app.component("MdEditor", MdEditor as any);
app.mount("#app");
