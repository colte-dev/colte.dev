<script context="module" lang="ts">
	import type { Load } from "@sveltejs/kit";

	export const load: Load = async ({ page }) => {
		return { props: { username: page.params.username } };
	};
</script>

<script lang="ts">
	import { onMount } from "svelte";
	import Link from "svelte-icons/md/MdLink.svelte";

	import { Post, Author, LinkIcon } from "$components";
	import { getUserDiscussions } from "$api";
	import type { Discussion } from "$api";

	export let username: string;
	export let posts: Discussion[] = [];
	export let author: Discussion["author"];

	onMount(async () => {
		posts = (await getUserDiscussions(username)).data;
		if (posts.length) author = posts[0].author;
	});
</script>

<svelte:head>
	<title>{username}</title>
</svelte:head>

{#if posts.length}
	<div class="container mx-auto p-4">
		<h1 class="py-7">
			<a href="https://github.com/colte-dev/forum/discussions">colte-dev/forum</a>
		</h1>

		<div class="grid grid-cols-4 gap-4">
			<div class="col-span-1 ">
				<div class="grid grid-cols-1 gap-y-4">
					<h4 class="mb-2">Postingan Oleh</h4>
					<Author avatarUrl={author.avatarUrl} name={author.login} label="" />
					<LinkIcon url={author.url} icon={Link} label="Buka di GitHub" />
				</div>
			</div>
			<div class="col-span-3">
				<div class="grid gap-y-4 lg:gap-y-6 grid-cols-1">
					{#each posts as post}
						<Post {post} />
					{/each}
				</div>
			</div>
		</div>
	</div>
{/if}
