import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import init from "helpers-wasm";

init()
	.then(() => {
		createApp(App).mount("#app");
	})
	.catch((err) => {
		console.log("error load the file:", err);
	});
