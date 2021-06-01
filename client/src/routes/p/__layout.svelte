<script lang="ts">
	import { goto } from "$app/navigation";

	import ArrowBack from "svelte-icons/md/MdArrowBack.svelte";
</script>

<div class="container mx-auto px-4 py-2 lg:py-4">
	<h4>
		<span on:click={() => goto("/")} class="cursor-pointer text-primary">
			<div class="w-6 inline-block align-middle">
				<ArrowBack />
			</div>
			<div class="inline-block align-middle">Home</div>
		</span>
	</h4>
</div>

<slot />

<style global lang="postcss">
	h2,
	h3,
	h4,
	h5,
	h6 {
		@apply my-2 lg:my-4;
	}

	h1 {
		font-size: 2rem;
		@apply my-6 lg:my-8 pb-6 lg:pb-8;
		@apply border-b;
	}

	h2 {
		font-size: 1.75rem;
	}

	h3 {
		font-size: 1.5rem;
	}

	h4 {
		font-size: 1.25rem;
	}

	.fab {
		padding-top: 12.5vh;
	}

	.post {
		max-width: 2048px;
	}

	.post-content {
		@apply px-4 md:px-8 py-2;
	}

	.post-content ul,
	ol {
		@apply px-8 py-4;
	}

	.highlight {
		background-color: #1a1c1f;
		@apply my-4 -mx-2;
		overflow-x: auto;
	}

	.highlight pre {
		@apply mx-4 my-2;
		display: inline-block;
	}

	.pl-c {
		color: #6a737d;
	}

	.pl-c1,
	.pl-s .pl-v {
		color: #2d76ca;
	}

	.pl-e,
	.pl-en {
		color: #9d7dda;
	}

	.pl-smi,
	.pl-s .pl-s1 {
		color: #cfcfcf;
	}

	.pl-ent {
		color: #22863a;
	}

	.pl-k {
		color: #d73a49;
	}

	.pl-s,
	.pl-pds,
	.pl-s .pl-pse .pl-s1,
	.pl-sr,
	.pl-sr .pl-cce,
	.pl-sr .pl-sre,
	.pl-sr .pl-sra {
		color: #96a3b3;
	}

	.pl-v,
	.pl-smw {
		color: #e36209;
	}

	.pl-bu {
		color: #b31d28;
	}

	.pl-ii {
		color: #fafbfc;
		background-color: #b31d28;
	}

	.pl-c2 {
		color: #fafbfc;
		background-color: #d73a49;
	}

	.pl-c2::before {
		content: "^M";
	}

	.pl-sr .pl-cce {
		font-weight: bold;
		color: #22863a;
	}

	.pl-ml {
		color: #735c0f;
	}

	.pl-mh,
	.pl-mh .pl-en,
	.pl-ms {
		font-weight: bold;
		color: #005cc5;
	}

	.pl-mi {
		font-style: italic;
		color: #24292e;
	}

	.pl-mb {
		font-weight: bold;
		color: #24292e;
	}

	.pl-md {
		color: #b31d28;
		background-color: #ffeef0;
	}

	.pl-mi1 {
		color: #22863a;
		background-color: #f0fff4;
	}

	.pl-mc {
		color: #e36209;
		background-color: #ffebda;
	}

	.pl-mi2 {
		color: #f6f8fa;
		background-color: #005cc5;
	}

	.pl-mdr {
		font-weight: bold;
		color: #6f42c1;
	}

	.pl-ba {
		color: #586069;
	}

	.pl-sg {
		color: #959da5;
	}

	.pl-corl {
		text-decoration: underline;
		color: #032f62;
	}
</style>
