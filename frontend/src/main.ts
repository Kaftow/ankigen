import { createApp } from "vue";
import App from "./App.vue";

import naive from "naive-ui";
import "./style.css";

import VMdEditor from "@kangc/v-md-editor";
import vuepressTheme from "@kangc/v-md-editor/lib/theme/vuepress.js";
import "@kangc/v-md-editor/lib/style/base-editor.css";
import "@kangc/v-md-editor/lib/theme/style/vuepress.css";

VMdEditor.use(vuepressTheme);

const app = createApp(App);
app.use(naive);
app.use(VMdEditor as any);
app.mount("#app");
