import { createApp } from "vue";
import App from "@/App.vue";
import router from "@/routers";
import { setupStore } from "@/stores/index";
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";
import naive from "naive-ui";

const app = createApp(App);
setupStore(app);
app.use(router);
app.use(naive);
app.component("MdEditor", MdEditor as any);
app.mount("#app");
