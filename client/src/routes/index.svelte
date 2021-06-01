<script lang="ts">
	import { onMount } from "svelte";
	import { NavMenuItem, Post } from "$components";
	import { getDiscussions } from "$api";
	import type { Discussion } from "$api";
	import Spinner from "$components/Spinner.svelte";

	let posts: Discussion[] = [];
	let postsPromise = getDiscussions().then(({ data }) => {
		posts = data;
	});

	const menu = [
		{
			text: "Sharing",
			icon: "💡",
		},
	];
</script>

<svelte:head>
	<title>Colte</title>
</svelte:head>

<div class="container mx-auto p-4">
	<h1 class="py-7">
		<a href="https://github.com/colte-dev/forum/discussions">colte-dev/forum</a>
	</h1>

	<div class="grid grid-cols-4 gap-4">
		<div class="col-span-1">
			{#each menu as m}
				<NavMenuItem emoji={m.icon} label={m.text} />
			{/each}
		</div>
		<div class="col-span-3">
			{#await postsPromise}
				<Spinner />
			{:then}
				{#each posts as post}
					<Post {post} />
				{/each}
			{/await}
		</div>
	</div>
</div>
