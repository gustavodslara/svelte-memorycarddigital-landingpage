<script>
	import "../app.css";
	import Navbar from "$lib/components/Navbar.svelte";
	import Footer from "$lib/components/Footer.svelte";
	import { afterNavigate } from "$app/navigation";
	import { page } from "$app/stores";
	import { tick, onMount } from "svelte";
	import "$lib/i18n"; // Initialize i18n
	import { waitLocale, isLoading } from "svelte-i18n";

	let { children } = $props();
	let isLocaleLoaded = $state(false);

	onMount(async () => {
		await waitLocale();
		isLocaleLoaded = true;
	});

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

{#if $isLoading || !isLocaleLoaded}
	<div class="loading-screen">Loading...</div>
{:else}
	<div class="app-wrapper">
		<Navbar />
		{#key $page.url.pathname}
			{@render children()}
		{/key}
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
