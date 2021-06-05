<script lang="ts">
	import MdThumbUp from "svelte-icons/md/MdThumbUp.svelte";
	import MdEmote from "svelte-icons/md/MdSentimentVerySatisfied.svelte";

	import Author from "./Author.svelte";
	import StatsButton from "./StatsButton.svelte";
	import type { Comment } from "src/api";

	export let comment: Comment;
</script>

<div class="rounded flex-col bg-accent w-full">
	<div class="p-4 grid grid-cols-1">
		<Author
			avatarUrl={comment.author.avatarUrl}
			name={comment.author.login}
			label={new Date(comment.createdAt).toLocaleString("id")}
			dense={true}
		/>

		<div class="py-2">
			{@html comment.bodyHTML}
		</div>

		<div class="align-middle grid grid-flow-col auto-cols-max gap-x-4 px-2 pt-4">
			<StatsButton icon={MdThumbUp} label={comment.upvoteCount} horizontal={true} iconSize={4} />
			<StatsButton
				icon={MdEmote}
				label={comment.reactions.totalCount}
				horizontal={true}
				iconSize={4}
			/>
		</div>
	</div>
</div>
