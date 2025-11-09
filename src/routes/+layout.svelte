<script>
	import '../app.css';
	import Navbar from '$lib/components/Navbar.svelte';
	import Footer from '$lib/components/Footer.svelte';
</script>

<div class="app-wrapper">
	<Navbar />
	<slot />
	<Footer />
</div>

<style>
	.app-wrapper {
		overflow-x: hidden;
	}
</style>
