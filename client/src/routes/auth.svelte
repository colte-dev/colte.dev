<script context="module" lang="ts">
	import type { Load } from "@sveltejs/kit";

	export const load: Load = async ({ page }) => {
		return { props: { code: page.query.get("code") } };
	};
</script>

<script lang="ts">
	import { onMount } from "svelte";
	import { goto } from "$app/navigation";
	import { authenticate } from "$api";

	export let code: string;

	onMount(async () => {
		const response = await authenticate(code);
		if (response.status === 204) await goto("/");
		else await goto("/login");
	});
</script>

<svelte:head>
	<title>Colte</title>
</svelte:head>

<div class="flex h-screen">
	<div class="m-auto text-primary">
		<h1>Authenticating...</h1>
	</div>
</div>
