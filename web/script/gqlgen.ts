// 通过 package.json 中的 script 调用,文件路径根目录为 web 文件夹。
import { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {

  generates: {
    "src/generated/workflow/": {
      preset: 'client',
      presetConfig: {
        gqlTagName: 'gql',
      },
      schema: ["../api/graphql/*.graphql"],
      documents: "src/service/**/*.ts",
    }
  },
  ignoreNoDocuments: true,
}


export default config
