import { defineConfig } from 'vite';
import { svelte } from '@sveltejs/vite-plugin-svelte';
import tailwindcss from '@tailwindcss/vite';
import path from 'path';

// https://vite.dev/config/
export default defineConfig({
	plugins: [svelte(), tailwindcss()],
	server: {
		cors: {
			origin: '*',
			methods: 'GET,HEAD,PUT,PATCH,POST,DELETE',
			preflightContinue: false,
			optionsSuccessStatus: 204
		},
		proxy: {
			'/api': { target: 'http://localhost:3000', changeOrigin: true }
		},
		allowedHosts: ['kocakhost']
	},
	resolve: {
		alias: {
			$lib: path.resolve('./src/lib')
		}
	}
});
