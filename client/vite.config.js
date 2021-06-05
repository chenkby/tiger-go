import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import styleImport from 'vite-plugin-style-import'

const { resolve } = require('path')

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'),
      vue: 'vue/dist/vue.esm-bundler.js', // @todo 对最终发布有什么影响 ？？
    },
  },
  // 引用elementUI全局scss变量文件,引用后自定义失效
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: '@import "@/style/var.scss";',
        // '@import "element-plus/packages/theme-chalk/src/common/var.scss";',
      },
    },
  },
  plugins: [
    vue(),
    styleImport({
      libs: [
        {
          libraryName: 'element-plus',
          esModule: true,
          ensureStyleFile: true,
          resolveStyle: (name) => {
            name = name.slice(3)
            return `element-plus/packages/theme-chalk/src/${name}.scss`
          },
          resolveComponent: (name) => `element-plus/lib/${name}`,
        },
      ],
    }),
  ],
})
