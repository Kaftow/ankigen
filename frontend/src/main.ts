import { createApp } from "vue";
import App from "./App.vue";
import VMdEditor from "@kangc/v-md-editor/lib/codemirror-editor";
import githubTheme from "@kangc/v-md-editor/lib/theme/github.js";
import * as Codemirror from "codemirror";
import hljs from "highlight.js";
import naive from "naive-ui";
import "./style.css";

const app = createApp(App);
VMdEditor.Codemirror = (Codemirror as any).default || Codemirror;
(VMdEditor as any).use(githubTheme, { Hljs: hljs });
app.use(VMdEditor as any);
app.use(naive);
app.mount("#app");
