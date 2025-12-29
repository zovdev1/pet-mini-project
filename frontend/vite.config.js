import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import legacy from '@vitejs/plugin-legacy'
// import { loadEnv } from 'vite'

// { mode }
export default defineConfig(() => {
  // const _env = loadEnv(mode, process.cwd(), '')

  return {
    plugins: [
      react(),
      legacy({
        targets: ['> 0.01%'],
      }),
    ],

    // define: {
    //   'import.meta.env.VITE_API_URL': JSON.stringify(env.VITE_API_URL || ''),
    // },
  }
})
