import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react-swc'
import { TanStackRouterVite } from '@tanstack/router-plugin/vite'
import {resolve} from 'path'

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [
		TanStackRouterVite(),
		react(),
	],
	resolve:{
		alias: {
			'@':resolve(__dirname,"src")
		}
	},
	server:{
		port: 3000,
	}
})
