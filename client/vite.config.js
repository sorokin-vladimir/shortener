import { defineConfig, loadEnv } from 'vite'

export default defineConfig(({mode}) => {
  const env = loadEnv(mode, process.cwd(), '')
  return {
    server: {
      port: parseInt(env.CLIENT_PORT, 10),
    },
  }
})
