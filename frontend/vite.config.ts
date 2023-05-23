import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'
import VueJsx from '@vitejs/plugin-vue-jsx'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

// https://vitejs.dev/config/
export default defineConfig({
  base: './',
  plugins: [
    vue(),
    VueJsx(),
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
  ],
  resolve: {
    alias: {
      '@': resolve(__dirname, './src'),
    },
    extensions: ['.ts', '.js', '.jsx', '.tsx']
  },
  css: {
    preprocessorOptions: {
      scss: {
        additionalData: '@import "./src/assets/styles/global.scss";@import "./src/assets/styles/common.scss";'
      }
    }
  },
  server: {
    host: '0.0.0.0',
    port: 8000,
  }
})
