<script lang="ts">
	import { SvelteComponent, createEventDispatcher } from "svelte";

	const dispatch = createEventDispatcher();

	export let label: string | number = "";
	export let icon: typeof SvelteComponent;
	export let horizontal: boolean = false;
	export let iconSize = 6;

	const className = horizontal ? "inline-block" : "";
</script>

<button on:click={() => dispatch("click")} class="focus:outline-none">
	<div class={`w-${iconSize} align-middle ${className}`}>
		<svelte:component this={icon} />
	</div>
	{#if label}
		<div class={`text-center ${className} ${horizontal ? "px-1" : ""}`}>
			{label}
		</div>
	{/if}
</button>
