<script lang="ts">
	import MdThumbUp from "svelte-icons/md/MdThumbUp.svelte";
	import MdComment from "svelte-icons/md/MdComment.svelte";
	import MdShare from "svelte-icons/md/MdShare.svelte";
	import MdEmote from "svelte-icons/md/MdSentimentVerySatisfied.svelte";
	import { goto } from "$app/navigation";

	import Author from "./Author.svelte";
	import StatsButton from "./StatsButton.svelte";
	import type { Discussion } from "src/api";
	import { activePost } from "$store/post";

	export let post: Discussion;

	const element = document.createElement("div");
	element.innerHTML = post.bodyHTML;
	const banner = element.querySelector("img");

	const clickHandler = () => {
		activePost.set(post);
		void goto(`/p/${post.number}`, {});
	};

	// export let post: Post;
</script>

<div class="rounded flex-col bg-accent w-full">
	{#if banner}
		<img
			class="banner mx-auto cursor-pointer"
			on:click={clickHandler}
			src={banner.src}
			alt={banner.alt}
		/>
	{/if}

	<div class="p-4 grid grid-cols-1">
		<h4 class="font-semibold cursor-pointer" on:click={clickHandler}>{post.title}</h4>

		<Author
			avatarUrl={post.author.avatarUrl}
			name={post.author.login}
			label={new Date(post.createdAt).toLocaleString()}
		/>

		<div class="align-middle grid grid-flow-col auto-cols-max gap-x-8 px-2 pt-4">
			<StatsButton icon={MdThumbUp} label={post.upvoteCount} horizontal={true} />
			<StatsButton icon={MdComment} label={post.comments.totalCount} horizontal={true} />
			<StatsButton icon={MdEmote} label={post.reactions.totalCount} horizontal={true} />
			<StatsButton icon={MdShare} horizontal={true} />
		</div>
	</div>
</div>

<style>
	.banner {
		max-height: 320px;
	}
</style>
