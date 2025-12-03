<script>
	import "../app.css";
	import Navbar from "$lib/components/Navbar.svelte";
	import Footer from "$lib/components/Footer.svelte";
	import { afterNavigate } from "$app/navigation";
	import { tick } from "svelte";
	import "$lib/i18n"; // Initialize i18n
	import { isLoading } from "svelte-i18n";

	afterNavigate(async () => {
		await tick();

		// Force immediate scroll to top without animation
		const html = document.documentElement;
		const originalScrollBehavior = html.style.scrollBehavior;
		html.style.scrollBehavior = "auto";

		window.scrollTo(0, 0);
		document.body.scrollTop = 0;
		html.scrollTop = 0;

		const scrollContainer = document.querySelector(".scroll-container");
		if (scrollContainer) {
			scrollContainer.scrollTop = 0;
		}

		// Restore original behavior
		setTimeout(() => {
			html.style.scrollBehavior = originalScrollBehavior;
		}, 10);
	});
</script>

{#if $isLoading}
	<div class="loading-screen">Loading...</div>
{:else}
	<div class="app-wrapper">
		<Navbar />
		<slot />
		<Footer />
	</div>
{/if}

<style>
	.app-wrapper {
		overflow-x: hidden;
	}

	.loading-screen {
		height: 100vh;
		width: 100vw;
		display: flex;
		align-items: center;
		justify-content: center;
		background: #0f172a;
		color: white;
		font-family: sans-serif;
	}
</style>
