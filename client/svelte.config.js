import preprocess from "svelte-preprocess";
import path from "path";
import adapt from "@sveltejs/adapter-static";

/** @type {import("@sveltejs/kit").Config} */
const config = {
	// Consult https://github.com/sveltejs/svelte-preprocess
	// for more information about preprocessors
	preprocess: [
		preprocess({
			postcss: true,
		}),
	],

	kit: {
		adapter: adapt({
			fallback: "index.html",
		}),
		files: {
			assets: "assets",
		},
		// hydrate the <div id="svelte"> element in src/app.html
		target: "#svelte",
		ssr: false,
		vite: {
			resolve: {
				alias: {
					$components: path.resolve("./src/components"),
					$api: path.resolve("./src/api"),
					$store: path.resolve("./src/store"),
				},
			},
		},
	},
};

export default config;
