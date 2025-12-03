<script>
	import { onMount } from "svelte";
	import { page } from "$app/stores";
	import { goto } from "$app/navigation";

	let scrolled = $state(false);
	let currentPath = $derived($page.url.pathname);

	// Determine if current page has light background
	let isLightPage = $derived(currentPath === "/features");

	// Determine if current page has gradient background (Download, Get Started)
	let isGradientPage = $derived(
		currentPath === "/download" || currentPath === "/get-started",
	);

	let logoCanvas;

	// Memory Card color scheme for logo (matching gradient text)
	const logoCardColor = {
		primary: "#6366f1",
		secondary: "#8b5cf6",
		accent: "#a78bfa",
	};

	/**
	 * @param {{ getContext: (arg0: string) => any; width: any; height: any; }} canvas
	 * @param {{ primary: any; secondary: any; accent: any; }} colorScheme
	 */
	function drawMemoryCard(canvas, colorScheme) {
		const ctx = canvas.getContext("2d");
		const width = canvas.width;
		const height = canvas.height;

		ctx.clearRect(0, 0, width, height);

		// Main body
		ctx.fillStyle = colorScheme.primary;
		ctx.fillRect(0, 0, width, height);

		// Top darker section (label area)
		ctx.fillStyle = colorScheme.secondary;
		ctx.fillRect(0, 0, width, height * 0.35);

		// Bottom accent strip
		ctx.fillStyle = colorScheme.accent;
		ctx.fillRect(0, height * 0.85, width, height * 0.15);

		// Connector notch at top
		ctx.fillStyle = colorScheme.secondary;
		const notchWidth = width * 0.6;
		const notchHeight = height * 0.08;
		const notchX = (width - notchWidth) / 2;
		ctx.fillRect(notchX, 0, notchWidth, notchHeight);

		// Side grip lines (left)
		ctx.strokeStyle = colorScheme.secondary;
		ctx.lineWidth = 1.2;
		for (let i = 0; i < 3; i++) {
			const y = height * 0.45 + i * 8;
			ctx.beginPath();
			ctx.moveTo(2, y);
			ctx.lineTo(8, y);
			ctx.stroke();
		}

		// Side grip lines (right)
		for (let i = 0; i < 3; i++) {
			const y = height * 0.45 + i * 8;
			ctx.beginPath();
			ctx.moveTo(width - 8, y);
			ctx.lineTo(width - 2, y);
			ctx.stroke();
		}

		// Minimalist icon circles (representing memory storage)
		ctx.fillStyle = colorScheme.accent;
		const iconY = height * 0.18;
		const spacing = width / 5;
		for (let i = 0; i < 3; i++) {
			ctx.beginPath();
			ctx.arc(spacing * (i + 1.5), iconY, 2.5, 0, Math.PI * 2);
			ctx.fill();
		}
	}

	onMount(async () => {
		const handleScroll = () => {
			scrolled = window.scrollY > 50;
		};

		window.addEventListener("scroll", handleScroll);

		// Draw memory card on canvas
		if (logoCanvas) {
			drawMemoryCard(logoCanvas, logoCardColor);
		}

		// Import GSAP for animations
		const { default: gsap } = await import("gsap");

		// Floating animation for the memory card
		gsap.to(".logo-memory-card", {
			y: -4,
			rotation: 2,
			duration: 2,
			repeat: -1,
			yoyo: true,
			ease: "sine.inOut",
		});

		return () => window.removeEventListener("scroll", handleScroll);
	});

	function navigate(path) {
		goto(path);
	}

	import SubscriptionModal from "./SubscriptionModal.svelte";
	let showModal = $state(false);
	let isMenuOpen = $state(false);

	function toggleMenu() {
		isMenuOpen = !isMenuOpen;
	}

	function closeMenu() {
		isMenuOpen = false;
	}

	import { auth, logout } from "$lib/stores/auth";

	let isAuthenticated = $state(false);
	let user = $state(null);

	auth.subscribe((value) => {
		isAuthenticated = value.isAuthenticated;
		user = value.user;
	});

	function handleLogout() {
		logout();
		goto("/");
	}

	import { t } from "svelte-i18n";
	import { base } from "$app/paths";
</script>

<SubscriptionModal isOpen={showModal} on:close={() => (showModal = false)} />

<nav
	class="navbar"
	class:scrolled
	class:light-page={isLightPage && !scrolled}
	class:gradient-page={isGradientPage && !scrolled}
>
	<div class="container">
		<a href="{base}/" data-sveltekit-preload-data class="logo">
			<div class="logo-memory-card">
				<canvas bind:this={logoCanvas} width="40" height="55"></canvas>
			</div>
			<span class="logo-text">memorycard.digital</span>
		</a>

		<ul class="nav-links">
			<li>
				<a href="{base}/" data-sveltekit-preload-data
					>{$t("navbar.home")}</a
				>
			</li>
			<li>
				<a href="{base}/about" data-sveltekit-preload-data
					>{$t("navbar.about")}</a
				>
			</li>
			<li>
				<a href="{base}/download" data-sveltekit-preload-data
					>{$t("navbar.download")}</a
				>
			</li>
		</ul>

		<div class="nav-actions">
			{#if isAuthenticated}
				<button
					class="subscribe-btn-nav"
					onclick={() => (showModal = true)}
				>
					{$t("navbar.subscribe")}
				</button>
				<a
					href="{base}/profile"
					class="profile-link"
					data-sveltekit-preload-data
				>
					<i class="fa-solid fa-user-circle"></i>
					<span>{$t("navbar.profile")}</span>
				</a>
				<button
					class="logout-btn"
					onclick={handleLogout}
					aria-label="Logout"
				>
					<i class="fa-solid fa-sign-out-alt"></i>
				</button>
			{:else}
				<a
					href="{base}/auth"
					class="subscribe-btn-nav"
					data-sveltekit-preload-data
				>
					{$t("navbar.signIn")}
				</a>
			{/if}
			<a
				href="{base}/get-started"
				data-sveltekit-preload-data
				class="cta-button">{$t("navbar.getStarted")}</a
			>
			<a
				href="https://github.com/gustavodslara/qt-memorycarddigital-app"
				target="_blank"
				rel="noopener noreferrer"
				class="github-button"
			>
				<i class="fab fa-github"></i>
				GitHub
			</a>
		</div>

		<button class="hamburger" onclick={toggleMenu} aria-label="Menu">
			<span class="bar" class:active={isMenuOpen}></span>
			<span class="bar" class:active={isMenuOpen}></span>
			<span class="bar" class:active={isMenuOpen}></span>
		</button>
	</div>

	<div class="mobile-menu" class:open={isMenuOpen}>
		<div class="mobile-nav-links">
			<a href="{base}/" data-sveltekit-preload-data onclick={closeMenu}
				>{$t("navbar.home")}</a
			>
			<a
				href="{base}/about"
				data-sveltekit-preload-data
				onclick={closeMenu}>{$t("navbar.about")}</a
			>
			<a
				href="{base}/download"
				data-sveltekit-preload-data
				onclick={closeMenu}>{$t("navbar.download")}</a
			>
			{#if isAuthenticated}
				<a
					href="{base}/profile"
					data-sveltekit-preload-data
					onclick={closeMenu}>{$t("navbar.profile")}</a
				>
				<button
					class="subscribe-btn-mobile"
					onclick={() => {
						showModal = true;
						closeMenu();
					}}
				>
					{$t("navbar.subscribe")}
				</button>
				<button
					class="subscribe-btn-mobile"
					style="background: transparent; border: 1px solid rgba(255,255,255,0.2); margin-top: 0.5rem;"
					onclick={() => {
						handleLogout();
						closeMenu();
					}}
				>
					{$t("navbar.logout")}
				</button>
			{:else}
				<a
					href="{base}/auth"
					class="subscribe-btn-mobile"
					data-sveltekit-preload-data
					onclick={closeMenu}
				>
					{$t("navbar.signIn")}
				</a>
			{/if}
			<a
				href="{base}/get-started"
				class="cta-button-mobile"
				data-sveltekit-preload-data
				onclick={closeMenu}>{$t("navbar.getStarted")}</a
			>
			<a
				href="https://github.com/gustavodslara/qt-memorycarddigital-app"
				target="_blank"
				rel="noopener noreferrer"
				class="github-button-mobile"
			>
				<i class="fab fa-github"></i>
				GitHub
			</a>
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

	/* Gradient page styling when not scrolled */
	.navbar.gradient-page {
		background: rgba(0, 0, 0, 0.1);
		backdrop-filter: blur(10px);
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
		display: flex;
		align-items: center;
		gap: 0.65rem;
		text-decoration: none;
		transition: transform 0.3s ease;
	}

	.logo:hover {
		transform: scale(1.05);
	}

	.logo:hover .logo-memory-card {
		filter: drop-shadow(0 8px 20px rgba(99, 102, 241, 0.6));
		transform: scale(1.15) rotate(-5deg);
	}

	.logo:hover .logo-text {
		transform: translateY(-2px);
		text-shadow: 0 4px 12px rgba(99, 102, 241, 0.5);
	}

	.logo-memory-card {
		cursor: pointer;
		filter: drop-shadow(0 4px 12px rgba(0, 0, 0, 0.3));
		transition: all 0.3s ease;
		will-change: transform;
	}

	.logo-memory-card canvas {
		display: block;
		border-radius: 3px;
	}

	.logo-text {
		font-size: 1.15rem;
		font-weight: 700;
		background: var(--gradient);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
		white-space: nowrap;
		transition: all 0.3s ease;
	}

	/* White text logo on gradient pages */
	.navbar.gradient-page .logo-text {
		background: white;
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
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
		content: "";
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
		justify-content: center;
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
		transition:
			transform 0.3s ease,
			box-shadow 0.3s ease;
		border: none;
		cursor: pointer;
		text-decoration: none;
		display: inline-block;
	}

	.cta-button:hover {
		transform: translateY(-2px);
		box-shadow: 0 10px 20px rgba(99, 102, 241, 0.3);
	}

	@media (max-width: 768px) {
		.nav-links {
			display: none;
		}

		.github-button {
			padding: 0.75rem 1rem;
		}

		.container {
			padding: 0 1rem;
		}

		.logo-memory-card canvas {
			width: 32px;
			height: 44px;
		}

		.logo-text {
			font-size: 0.95rem;
		}

		.logo {
			gap: 0.5rem;
		}

		.subscribe-btn-nav {
			padding: 0.5rem 1rem;
			font-size: 0.9rem;
		}
	}

	.subscribe-btn-nav {
		background: rgba(255, 255, 255, 0.1);
		color: var(--text-light);
		border: 1px solid rgba(255, 255, 255, 0.2);
		padding: 0.75rem 1.5rem;
		border-radius: 50px;
		font-weight: 600;
		cursor: pointer;
		transition: all 0.3s ease;
		backdrop-filter: blur(5px);
	}

	.subscribe-btn-nav:hover {
		background: rgba(255, 255, 255, 0.2);
		border-color: rgba(255, 255, 255, 0.4);
		transform: translateY(-2px);
	}

	.navbar.light-page .subscribe-btn-nav {
		background: rgba(0, 0, 0, 0.05);
		color: var(--text-dark);
		border-color: rgba(0, 0, 0, 0.1);
	}

	.navbar.light-page .subscribe-btn-nav:hover {
		background: rgba(0, 0, 0, 0.1);
		border-color: rgba(0, 0, 0, 0.2);
	}

	/* Hamburger Menu */
	.hamburger {
		display: none;
		flex-direction: column;
		justify-content: space-between;
		width: 30px;
		height: 21px;
		background: transparent;
		border: none;
		cursor: pointer;
		z-index: 1001;
		padding: 0;
	}

	.bar {
		width: 100%;
		height: 3px;
		background-color: var(--text-light);
		border-radius: 3px;
		transition: all 0.3s ease;
	}

	.navbar.light-page .bar {
		background-color: var(--text-dark);
	}

	.bar.active:nth-child(1) {
		transform: translateY(9px) rotate(45deg);
	}

	.bar.active:nth-child(2) {
		opacity: 0;
	}

	.bar.active:nth-child(3) {
		transform: translateY(-9px) rotate(-45deg);
	}

	/* Mobile Menu */
	.mobile-menu {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100vh;
		background: rgba(15, 23, 42, 0.98);
		backdrop-filter: blur(20px);
		z-index: 1000;
		display: flex;
		align-items: center;
		justify-content: center;
		opacity: 0;
		pointer-events: none;
		transition: opacity 0.3s ease;
	}

	.mobile-menu.open {
		opacity: 1;
		pointer-events: all;
	}

	.mobile-nav-links {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 2rem;
		width: 100%;
		padding: 2rem;
	}

	.mobile-nav-links a {
		color: var(--text-light);
		font-size: 1.5rem;
		font-weight: 600;
		text-decoration: none;
		transition: color 0.3s ease;
	}

	.mobile-nav-links a:hover {
		color: var(--primary-color);
	}

	.subscribe-btn-mobile {
		background: var(--gradient);
		color: white;
		border: none;
		padding: 1rem 2rem;
		border-radius: 50px;
		font-size: 1.2rem;
		font-weight: 600;
		width: 100%;
		max-width: 300px;
		cursor: pointer;
	}

	.cta-button-mobile {
		background: rgba(255, 255, 255, 0.1);
		border: 1px solid rgba(255, 255, 255, 0.2);
		color: white;
		padding: 1rem 2rem;
		border-radius: 50px;
		font-size: 1.2rem;
		font-weight: 600;
		width: 100%;
		max-width: 300px;
		text-align: center;
	}

	.github-button-mobile {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 0.5rem;
		color: #94a3b8;
		font-size: 1.1rem;
		text-decoration: none;
	}

	@media (max-width: 768px) {
		.hamburger {
			display: flex;
		}

		.nav-actions {
			display: none;
		}
	}

	.profile-link {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		color: var(--text-light);
		text-decoration: none;
		font-weight: 500;
		transition: color 0.3s ease;
	}

	.profile-link:hover {
		color: var(--primary-color);
	}

	.navbar.light-page .profile-link {
		color: var(--text-dark);
	}

	.logout-btn {
		background: none;
		border: none;
		color: var(--text-light);
		cursor: pointer;
		font-size: 1.1rem;
		transition: color 0.3s ease;
		padding: 0.5rem;
	}

	.logout-btn:hover {
		color: #ef4444;
	}

	.navbar.light-page .logout-btn {
		color: var(--text-dark);
	}
</style>
