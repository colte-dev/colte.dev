<script lang="ts">
	import { goto } from "$app/navigation";

	export let name: string;
	export let avatarUrl: string;
	export let label: string;
</script>

<div on:click={() => goto("/" + name)} class="cursor-pointer">
	<div class="flex items-center ">
		<img class="rounded-full" width="44" height="44" src={avatarUrl} alt="avatar" />
		<div class="pl-4 py-2">
			<div class="text-lg text-gray-300">{name}</div>
			<div class="text-xs text-gray-400 ">
				{label}
			</div>
		</div>
	</div>
</div>
