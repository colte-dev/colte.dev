<script lang="ts">
	export let size: number = 16;
	export let border: number = 4;

	const classes = `h-${size} w-${size} border-${border} border-t-${border}`;
</script>

<div class={`loader ease-linear rounded-full mx-auto border-gray-200 ${classes}`} />

<style>
	.loader {
		border-top-color: #2f99ff;
		-webkit-animation: spinner 1.5s linear infinite;
		animation: spinner 1.5s linear infinite;
	}

	@-webkit-keyframes spinner {
		0% {
			-webkit-transform: rotate(0deg);
		}
		100% {
			-webkit-transform: rotate(360deg);
		}
	}

	@keyframes spinner {
		0% {
			transform: rotate(0deg);
		}
		100% {
			transform: rotate(360deg);
		}
	}
</style>
