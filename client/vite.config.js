import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import styleImport from 'vite-plugin-style-import'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    host: '0.0.0.0',
    open: '/',
    proxy: {
      // 字符串简写写法
      '/foo': 'http://localhost:4567/foo',
      // 选项写法
      '/api': {
        target: 'http://localhost:8081/v1/',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      },
      // 正则表达式写法
      '^/fallback/.*': {
        target: 'http://jsonplaceholder.typicode.com',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/fallback/, '')
      }
    }
  },
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src')
      //vue: 'vue/dist/vue.esm-bundler.js' // @todo 对最终发布有什么影响 ？？
    }
  },
  build: {
    sourcemap: true, // 构建后是否生成 source map 文件。
    cssCodeSplit: false // 启用/禁用 CSS 代码拆分。如果禁用，整个项目中的所有 CSS 将被提取到一个 CSS 文件中。
  },
  // 引用elementUI全局scss变量文件,引用后自定义失效
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: '@import "@/style/var.scss";'
        // '@import "element-plus/packages/theme-chalk/src/common/var.scss";',
      }
    }
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
          resolveComponent: (name) => `element-plus/lib/${name}`
        }
      ]
    })
  ]
})
