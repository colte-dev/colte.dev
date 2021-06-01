<script lang="ts">
	import type { SvelteComponent } from "svelte";

	export let url: string;
	export let label: string;
	export let icon: typeof SvelteComponent;
</script>

<a href={url}>
	<div class="w-6 inline-block align-middle text-gray-400">
		<svelte:component this={icon} />
	</div>
	{label}
</a>
