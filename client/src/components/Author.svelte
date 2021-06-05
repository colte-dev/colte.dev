<script lang="ts">
	import { goto } from "$app/navigation";

	export let name: string;
	export let avatarUrl: string;
	export let label: string;
	export let dense = false;

	const size = !dense ? 44 : 32;
	const nameClasses = dense ? "inline-block mr-1 text-md" : "text-lg";
</script>

<div on:click={() => goto("/" + name)} class="cursor-pointer">
	<div class="flex items-center ">
		<img class="rounded-full" width={size} height={size} src={avatarUrl} alt="avatar" />
		<div class="pl-4 py-2">
			<div class={`text-gray-300 ${nameClasses}`}>{name}</div>
			<div class={`text-xs text-gray-400 ${dense ? "inline-block" : ""}`}>{label}</div>
		</div>
	</div>
</div>
