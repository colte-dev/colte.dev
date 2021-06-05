<script context="module" lang="ts">
	import type { Load } from "@sveltejs/kit";

	export const load: Load = async ({ page }) => {
		return { props: { slug: page.params.slug } };
	};
</script>

<script lang="ts">
	import { onMount, tick } from "svelte";
	import { goto } from "$app/navigation";
	import ThumbUp from "svelte-icons/md/MdThumbUp.svelte";
	import Comment from "svelte-icons/md/MdComment.svelte";
	import Share from "svelte-icons/md/MdShare.svelte";
	import Link from "svelte-icons/md/MdLink.svelte";
	import Emote from "svelte-icons/md/MdSentimentVerySatisfied.svelte";

	import { Author, StatsButton, LinkIcon, HeaderList } from "$components";
	import { getDiscussionByNumber } from "$api";
	import type { Discussion } from "$api";
	import { activePost } from "$store/post";
	import slugify from "slugify";

	interface Header {
		level: number;
		label: string;
		encoded: string;
		id: string;
	}

	export let slug: string;
	const number = slug.split("-").pop();

	let post: Discussion | null = null;
	let headers: Header[] = [];
	const asideClass = "hidden lg:block h-screen sticky top-0 px-4 xl:px-8 2xl:w-1/4";

	const scrollTo = (id: string) => document.getElementById(id).scrollIntoView();

	const fetchPost = async () => {
		if ($activePost?.number === +number) post = $activePost;
		else {
			const response = await getDiscussionByNumber(+number);
			if (response.status === 200) post = response.data;
		}
	};

	const getHeaders = (html: string): { headers: Header[]; html: string } => {
		const headers = [];
		const element = document.createElement("div");
		element.innerHTML = html;

		const nodeHeaders: Element[] = Array.apply(
			null,
			element.querySelectorAll("h2, h3, h4, h5, h6")
		);

		nodeHeaders.forEach((h) => {
			const encoded = slugify(h.textContent, { lower: true });
			const existingIds = headers.filter((h) => h.encoded === encoded).length;
			const id = encoded + (existingIds > 0 ? "-" + existingIds : "");
			h.setAttribute("id", id);
			headers.push({
				level: +h.tagName.replace("H", ""),
				label: h.textContent,
				encoded,
				id,
			});
		});
		const smallestHeaderLevel = [...headers].sort((a, b) => a.level - b.level).shift().level;
		headers.forEach((h) => (h.level -= smallestHeaderLevel));
		return {
			headers,
			html: element.innerHTML,
		};
	};

	const updatePath = () => {
		const path = slugify(post.title, { lower: true }) + "-9" + window.location.hash;
		window.history.replaceState("", "", path);
	};

	onMount(async () => {
		await fetchPost();
		if (!post) return await goto("/");

		const { headers: h, html } = getHeaders(post.bodyHTML);
		post.bodyHTML = html;
		headers = h;

		updatePath();

		await tick();
		if (window.location.hash) scrollTo(window.location.hash.substring(1));
	});
</script>

<svelte:head>
	<title>{post?.title || "Colte"}</title>
</svelte:head>

{#if post}
	<div class="flex post mx-auto">
		<aside class={asideClass + " 2xl:pr-16 fab"}>
			<div class="grid grid-cols-1 gap-y-8 float-right w-6">
				<StatsButton icon={ThumbUp} label={post.upvoteCount} />
				<StatsButton
					on:click={() => scrollTo("comments")}
					icon={Comment}
					label={post.comments.totalCount}
				/>
				<StatsButton icon={Emote} label={post.reactions.totalCount} />
				<StatsButton icon={Share} />
			</div>
		</aside>

		<main class="w-full lg:w-3/5 2xl:w-1/2 mb-32">
			<div class="bg-accent post-content pb-6">
				{@html post.bodyHTML}
			</div>

			<!-- <div id="comments" class="bg-accent w-full my-8 p-4 md:px-8" /> -->
		</main>

		<aside class={asideClass + " lg:w-1/3 max-w-sm overflow-y-auto sticky pt-12"}>
			<div
				class="uppercase font-semibold mb-3 text-sm cursor-pointer text-gray-300"
				on:click={() => window.scrollTo(0, 0)}
			>
				{post.title}
			</div>

			<ul class="overflow-x-hidden">
				{#each headers as h}
					<HeaderList label={h.label} level={h.level} target={h.id} />
				{/each}

				<div class="mt-3">
					<HeaderList target="comments" label="Komentar" />
				</div>
			</ul>

			<div class="border-t border-gray-500 my-8" />

			<div class="my-6">
				<LinkIcon url={post.url} icon={Link} label="Buka di GitHub" />
			</div>

			<Author
				avatarUrl={post.author.avatarUrl}
				name={post.author.login}
				label={new Date(post.createdAt).toLocaleString()}
			/>
		</aside>
	</div>

	<footer class="lg:hidden fixed bottom-0 w-full h-12 bg-gray-800">
		<div class="grid grid-flow-col h-full text-center align-middle">
			<StatsButton icon={ThumbUp} label={post.upvoteCount} horizontal={true} />
			<StatsButton
				on:click={() => scrollTo("comments")}
				icon={Comment}
				label={post.comments.totalCount}
				horizontal={true}
			/>
			<StatsButton icon={Emote} label={post.reactions.totalCount} horizontal={true} />
			<StatsButton icon={Share} horizontal={true} />
		</div>
	</footer>
{/if}
