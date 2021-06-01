const colors = require("tailwindcss/colors");

module.exports = {
	mode: "jit",
	purge: {
		enabled: false,
		content: ["./src/**/*.{html,js,svelte,ts}"],
	},
	corePlugins: {
		listStyleType: false,
	},
	theme: {
		colors: {
			primary: {
				DEFAULT: "#2F99FF",
				50: "#8BC6FF",
				100: "#81C1FF",
				200: "#6CB7FF",
				300: "#58ADFF",
				400: "#43A3FF",
				500: "#2F99FF",
				600: "#0080FB",
				700: "#0066C8",
				800: "#004C95",
				900: "#003262",
			},
			accent: "#0D1117",

			transparent: "transparent",
			current: "currentColor",

			black: colors.black,
			white: colors.white,
			gray: colors.coolGray,
			red: colors.red,
			yellow: colors.amber,
			green: colors.emerald,
			blue: colors.blue,
			indigo: colors.indigo,
			purple: colors.violet,
			pink: colors.pink,
		},
	},
	plugins: [],
};
