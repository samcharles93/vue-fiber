import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      "/api/v1": {
        target: "http://localhost:3000",
        changeOrigin: true,
        rewrite: (path) => path.replace("\^\/api/v1", "")
      }
    }
  }
})
