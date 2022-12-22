import { defineConfig } from "vite";
import wasm_pack from "vite-plugin-wasm-pack";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [wasm_pack(["./helpers-wasm"]), vue()],
});
