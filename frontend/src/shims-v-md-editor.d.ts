declare module "@kangc/v-md-editor/lib/codemirror-editor" {
  import type { Plugin, DefineComponent } from "vue";
  const VMdEditor: Plugin &
    DefineComponent & {
      Codemirror?: any;
      use?: (theme: any, options?: any) => void;
    };
  export default VMdEditor;
}

declare module "@kangc/v-md-editor/*" {
  const mod: any;
  export default mod;
}

declare module "codemirror";

declare module "@kangc/v-md-editor" {
  import { Plugin, DefineComponent } from "vue";

  const VMdEditor: Plugin &
    DefineComponent & {
      use?: (theme: any, options?: any) => void;
    };
  export default VMdEditor;
}

declare module "@kangc/v-md-editor/lib/theme/vuepress.js" {
  const vuepressTheme: any;
  export default vuepressTheme;
}
