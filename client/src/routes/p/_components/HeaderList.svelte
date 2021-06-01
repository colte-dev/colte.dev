<script lang="ts">
	export let level: number = 0;
	export let target: string;
	export let label: string;
</script>

<li class="block duration-200 py-1" style={`padding-left: ${level * 1}rem`}>
	<a href={`#${target}`}>
		{label}
	</a>
</li>
