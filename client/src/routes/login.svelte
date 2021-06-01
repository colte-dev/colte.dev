<script lang="ts">
	import { goto } from "$app/navigation";
	import { checkAuth } from "$api";

	const checkAuthPromise = checkAuth().then(async ({ status }) => {
		if (status === 204) await goto("/");
	});

	const getCurrentHost = () => {
		return window.location.protocol + "//" + window.location.host + "/auth";
	};

	const loginUrl =
		"https://github.com/login/oauth/authorize?client_id=Iv1.0ff6f861106ed6a3&redirect_uri=" +
		getCurrentHost();
</script>

{#await checkAuthPromise then _}
	<div class="flex h-screen">
		<div class="m-auto">
			<div class="text-center text-xl lg:text-3xl">
				Harap Login Menggunakan GitHub untuk Menggunakan Aplikasi Ini
			</div>
			<div class="text-center">
				<a
					href={loginUrl}
					class="rounded-sm px-3 py-2 cursor-pointer bg-green-500 hover:bg-green-600 text-white my-12"
					type="submit"
				>
					Login dengan GitHub
				</a>
			</div>
		</div>
	</div>
{/await}
