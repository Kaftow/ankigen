declare module "@kangc/v-md-editor" {
  import { App } from "vue";

  interface VMdEditor {
    use(theme: any, options?: any): void;
  }

  const VMdEditor: VMdEditor;
  export default VMdEditor;
}

declare module "@kangc/v-md-editor/lib/theme/vuepress.js" {
  const vuepressTheme: any;
  export default vuepressTheme;
}
