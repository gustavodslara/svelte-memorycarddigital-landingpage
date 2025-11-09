<script>
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';

	let scrolled = $state(false);
	let currentPath = $derived($page.url.pathname);
	
	// Determine if current page has light background
	let isLightPage = $derived(currentPath === '/features');

	onMount(() => {
		const handleScroll = () => {
			scrolled = window.scrollY > 50;
		};

		window.addEventListener('scroll', handleScroll);
		return () => window.removeEventListener('scroll', handleScroll);
	});

	function navigate(path) {
		goto(path);
	}
</script>

<nav class="navbar" class:scrolled class:light-page={isLightPage && !scrolled}>
	<div class="container">
		<a href="/" data-sveltekit-preload-data class="logo">
			<span class="logo-text">Memory Card Digital</span>
		</a>
		
		<ul class="nav-links">
			<li><a href="/features" data-sveltekit-preload-data>Features</a></li>
			<li><a href="/about" data-sveltekit-preload-data>About</a></li>
			<li><a href="/contact" data-sveltekit-preload-data>Contact</a></li>
		</ul>
		
		<div class="nav-actions">
			<a href="https://github.com/gustavodslara/svelte-memorycarddigital-landingpage" target="_blank" rel="noopener noreferrer" class="github-button">
				<i class="fab fa-github"></i>
				GitHub
			</a>
			<button class="cta-button">Get Started</button>
		</div>
	</div>
</nav>

<style>
	.navbar {
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		z-index: 1000;
		padding: 1.5rem 0;
		transition: all 0.3s ease;
		background: transparent;
	}

	.navbar.scrolled {
		background: rgba(15, 23, 42, 0.95);
		backdrop-filter: blur(10px);
		padding: 1rem 0;
		box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
	}

	/* Light page styling when not scrolled */
	.navbar.light-page {
		background: rgba(255, 255, 255, 0.9);
		backdrop-filter: blur(10px);
		box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
	}

	.container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 0 2rem;
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.logo {
		font-size: 1.5rem;
		font-weight: 700;
		background: var(--gradient);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
		transition: transform 0.3s ease;
		text-decoration: none;
	}

	.logo:hover {
		transform: scale(1.05);
	}

	.nav-links {
		display: flex;
		list-style: none;
		gap: 2rem;
	}

	.nav-links a {
		color: var(--text-light);
		font-weight: 500;
		transition: color 0.3s ease;
		position: relative;
		text-decoration: none;
	}

	/* Dark text on light pages */
	.navbar.light-page .nav-links a {
		color: var(--text-dark);
	}

	.nav-links a::after {
		content: '';
		position: absolute;
		bottom: -5px;
		left: 0;
		width: 0;
		height: 2px;
		background: var(--gradient);
		transition: width 0.3s ease;
	}

	.nav-links a:hover::after {
		width: 100%;
	}

	.nav-actions {
		display: flex;
		align-items: center;
		gap: 1rem;
	}

	.github-button {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		background: transparent;
		color: var(--text-light);
		padding: 0.75rem 1.5rem;
		border-radius: 50px;
		font-weight: 600;
		border: 2px solid rgba(255, 255, 255, 0.3);
		transition: all 0.3s ease;
		text-decoration: none;
		font-size: 1rem;
	}

	.navbar.light-page .github-button {
		color: var(--text-dark);
		border-color: rgba(0, 0, 0, 0.2);
	}

	.github-button:hover {
		background: rgba(255, 255, 255, 0.1);
		border-color: rgba(255, 255, 255, 0.5);
		transform: translateY(-2px);
	}

	.navbar.light-page .github-button:hover {
		background: rgba(0, 0, 0, 0.05);
		border-color: rgba(0, 0, 0, 0.3);
	}

	.github-button i {
		font-size: 1.1rem;
	}

	.cta-button {
		background: var(--gradient);
		color: white;
		padding: 0.75rem 1.5rem;
		border-radius: 50px;
		font-weight: 600;
		transition: transform 0.3s ease, box-shadow 0.3s ease;
		border: none;
		cursor: pointer;
	}

	.cta-button:hover {
		transform: translateY(-2px);
		box-shadow: 0 10px 20px rgba(99, 102, 241, 0.3);
	}

	@media (max-width: 768px) {
		.nav-links {
			display: none;
		}

		.github-button span {
			display: none;
		}

		.github-button {
			padding: 0.75rem 1rem;
		}
		
		.container {
			padding: 0 1rem;
		}
	}
</style>
