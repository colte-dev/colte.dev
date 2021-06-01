<script lang="ts">
	import { logout } from "$api";
	import { goto } from "$app/navigation";

	logout().then(() => {
		goto("/");
	});
</script>
