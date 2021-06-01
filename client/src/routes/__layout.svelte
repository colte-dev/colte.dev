<script lang="ts">
	import { checkAuth } from "$api";
	import { goto } from "$app/navigation";

	const authPromise = checkAuth().then((response) => {
		if (response.status != 204) goto("/login");
	});
</script>

{#await authPromise then _}
	<slot />
{/await}

<style global lang="postcss">
	@tailwind base;
	@tailwind components;
	@tailwind utilities;

	html {
		scroll-behavior: smooth;
	}

	a {
		@apply text-primary;
	}

	ul {
		list-style: disc;
	}

	ol {
		list-style: decimal;
	}

	body {
		background-color: #090c10;
		color: #c9d1d9;
		font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen-Sans, Ubuntu,
			Cantarell, "Helvetica Neue", sans-serif;
	}

	h1 {
		font-size: 3rem;
		line-height: 1;
	}
	h2 {
		font-size: 2.25rem;
		line-height: 2.5rem;
	}
	h3 {
		font-size: 1.875rem;
		line-height: 2.25rem;
	}
	h4 {
		font-size: 1.5rem;
		line-height: 2rem;
	}
	h5 {
		font-size: 1.25rem;
		line-height: 1.75rem;
	}
</style>
