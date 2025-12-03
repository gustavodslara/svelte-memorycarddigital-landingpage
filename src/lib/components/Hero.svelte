<script>
	import { onMount, tick } from "svelte";
	import { t } from "svelte-i18n";

	/**
	 * @type {any[]}
	 */
	let canvasRefs = [];
	/**
	 * @type {HTMLElement}
	 */
	let scrollContainer;

	// PS2 Memory Card colors (minimalist solid colors)
	const memoryCardColors = [
		{ primary: "#1e40af", secondary: "#1e3a8a", accent: "#3b82f6" }, // Blue
		{ primary: "#7c2d12", secondary: "#991b1b", accent: "#dc2626" }, // Red
		{ primary: "#065f46", secondary: "#064e3b", accent: "#10b981" }, // Green
		{ primary: "#1f2937", secondary: "#111827", accent: "#4b5563" }, // Dark Gray
		{ primary: "#7e22ce", secondary: "#581c87", accent: "#a855f7" }, // Purple
	];

	// Content for different scroll sections
	// Content for different scroll sections
	let sections = $derived([
		{
			title: $t("hero.backup.title"),
			highlight: $t("hero.backup.highlight"),
			subtitle: $t("hero.backup.subtitle"),
			buttons: [$t("hero.backup.button1"), $t("hero.backup.button2")],
		},
		{
			title: $t("hero.versioning.title"),
			highlight: $t("hero.versioning.highlight"),
			subtitle: $t("hero.versioning.subtitle"),
			buttons: [
				$t("hero.versioning.button1"),
				$t("hero.versioning.button2"),
			],
		},
		{
			title: $t("hero.cloud.title"),
			highlight: $t("hero.cloud.highlight"),
			subtitle: $t("hero.cloud.subtitle"),
			buttons: [$t("hero.cloud.button1"), $t("hero.cloud.button2")],
		},
		{
			title: $t("hero.security.title"),
			highlight: $t("hero.security.highlight"),
			subtitle: $t("hero.security.subtitle"),
			buttons: [$t("hero.security.button1"), $t("hero.security.button2")],
		},
		{
			title: $t("hero.download.title"),
			highlight: $t("hero.download.highlight"),
			subtitle: $t("hero.download.subtitle"),
			buttons: [$t("hero.download.button1"), $t("hero.download.button2")],
		},
	]);

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
		ctx.lineWidth = 2;
		for (let i = 0; i < 5; i++) {
			const y = height * 0.45 + i * 15;
			ctx.beginPath();
			ctx.moveTo(5, y);
			ctx.lineTo(15, y);
			ctx.stroke();
		}

		// Side grip lines (right)
		for (let i = 0; i < 5; i++) {
			const y = height * 0.45 + i * 15;
			ctx.beginPath();
			ctx.moveTo(width - 15, y);
			ctx.lineTo(width - 5, y);
			ctx.stroke();
		}

		// Minimalist icon circles (representing memory storage)
		ctx.fillStyle = colorScheme.accent;
		const iconY = height * 0.18;
		const spacing = width / 5;
		for (let i = 0; i < 3; i++) {
			ctx.beginPath();
			ctx.arc(spacing * (i + 1.5), iconY, 4, 0, Math.PI * 2);
			ctx.fill();
		}
	}

	onMount(() => {
		/** @type {gsap.Context | null} */
		let ctx = null;
		/** @type {typeof import("gsap/ScrollTrigger").ScrollTrigger | null} */
		let ScrollTriggerInstance = null;

		const init = async () => {
			await tick();
			if (!scrollContainer) return;

			const { default: gsap } = await import("gsap");
			const { ScrollTrigger } = await import("gsap/ScrollTrigger");
			ScrollTriggerInstance = ScrollTrigger;

			gsap.registerPlugin(ScrollTrigger);

			// Draw memory cards on canvas
			canvasRefs.forEach((canvas, index) => {
				if (canvas) {
					drawMemoryCard(canvas, memoryCardColors[index]);
				}
			});

			ctx = gsap.context(() => {
				// 1. Floating Animation (Applied to inner .memory-card)
				// This runs independently of scroll
				const cards = document.querySelectorAll(".memory-card");
				cards.forEach((card, index) => {
					gsap.to(card, {
						y: -20, // Float up
						rotation: gsap.utils.random(-5, 5),
						duration: gsap.utils.random(2, 3),
						repeat: -1,
						yoyo: true,
						ease: "sine.inOut",
						delay: index * 0.2,
					});
				});

				// 2. Scroll Animation (Applied to outer .card-wrapper)
				// This controls the position on screen

				// Pin the entire hero section
				ScrollTrigger.create({
					trigger: scrollContainer,
					start: "top top",
					end: "+=400%",
					pin: ".hero",
					pinSpacing: false,
					scrub: true,
					anticipatePin: 1,
					fastScrollEnd: true,
				});

				// Master timeline for scroll movements
				const masterTimeline = gsap.timeline({
					defaults: { overwrite: "auto" }, // Prevent conflicts
					scrollTrigger: {
						trigger: scrollContainer,
						start: "top top",
						end: "+=400%",
						scrub: 1.5,
						invalidateOnRefresh: true,
						fastScrollEnd: true,
					},
				});

				// Initial State (implicitly defined by CSS, but good to know):
				// All cards are centered at (0,0) in the stack

				// --- Section 0 to 1: Chaotic Explosion (Transform -> Store) ---
				masterTimeline
					.to(
						".card-wrapper-1",
						{
							x: -400,
							y: -300,
							rotation: -45,
							scale: 0.8,
							duration: 1,
						},
						0,
					)
					.to(
						".card-wrapper-2",
						{
							x: 400,
							y: -200,
							rotation: 60,
							scale: 1.2,
							duration: 1,
						},
						0,
					)
					.to(
						".card-wrapper-3",
						{
							x: -300,
							y: 300,
							rotation: -90,
							scale: 0.9,
							duration: 1,
						},
						0,
					)
					.to(
						".card-wrapper-4",
						{
							x: 350,
							y: 250,
							rotation: 120,
							scale: 1.1,
							duration: 1,
						},
						0,
					)
					.to(
						".card-wrapper-5",
						{
							x: 0,
							y: -400,
							rotation: 180,
							scale: 0.7,
							duration: 1,
						},
						0,
					)

					// Content transition
					.to(
						".section-0",
						{ autoAlpha: 0, y: -50, duration: 0.5 },
						0.2,
					)
					.to(
						".section-1",
						{ autoAlpha: 1, y: 0, duration: 0.5 },
						0.5,
					);

				// --- Section 1 to 2: Swirl/Orbit (Store -> Share) ---
				masterTimeline
					.to(
						".card-wrapper-1",
						{
							x: 300,
							y: -300,
							rotation: 45,
							scale: 1.1,
							duration: 1,
						},
						1,
					)
					.to(
						".card-wrapper-2",
						{
							x: 300,
							y: 300,
							rotation: 135,
							scale: 0.8,
							duration: 1,
						},
						1,
					)
					.to(
						".card-wrapper-3",
						{
							x: -300,
							y: 300,
							rotation: 225,
							scale: 1.2,
							duration: 1,
						},
						1,
					)
					.to(
						".card-wrapper-4",
						{
							x: -300,
							y: -300,
							rotation: 315,
							scale: 0.9,
							duration: 1,
						},
						1,
					)
					.to(
						".card-wrapper-5",
						{
							x: 0,
							y: 0,
							rotation: 360,
							scale: 1.5,
							zIndex: 20,
							duration: 1,
						},
						1,
					)

					// Content transition
					.to(
						".section-1",
						{ autoAlpha: 0, y: -50, duration: 0.5 },
						1.2,
					)
					.to(
						".section-2",
						{ autoAlpha: 1, y: 0, duration: 0.5 },
						1.5,
					);

				// --- Section 2 to 3: Cross Over (Share -> Preserve) ---
				masterTimeline
					.to(
						".card-wrapper-1",
						{
							x: -500,
							y: 100,
							rotation: -90,
							scale: 0.7,
							duration: 1,
						},
						2,
					)
					.to(
						".card-wrapper-2",
						{
							x: -200,
							y: -400,
							rotation: 0,
							scale: 1,
							duration: 1,
						},
						2,
					)
					.to(
						".card-wrapper-3",
						{
							x: 200,
							y: 400,
							rotation: 180,
							scale: 1,
							duration: 1,
						},
						2,
					)
					.to(
						".card-wrapper-4",
						{
							x: 500,
							y: -100,
							rotation: 90,
							scale: 0.7,
							duration: 1,
						},
						2,
					)
					.to(
						".card-wrapper-5",
						{
							x: 0,
							y: -500,
							rotation: 720,
							scale: 0.5,
							duration: 1,
						},
						2,
					)

					// Content transition
					.to(
						".section-2",
						{ autoAlpha: 0, y: -50, duration: 0.5 },
						2.2,
					)
					.to(
						".section-3",
						{ autoAlpha: 1, y: 0, duration: 0.5 },
						2.5,
					);

				// --- Section 3 to 4: Final Assembly (Preserve -> Download) ---
				// Bring them all back to a neat row or fan for the download section
				masterTimeline
					.to(
						".card-wrapper-1",
						{ x: -240, y: 0, rotation: -10, scale: 1, duration: 1 },
						3,
					)
					.to(
						".card-wrapper-2",
						{ x: -120, y: 0, rotation: -5, scale: 1, duration: 1 },
						3,
					)
					.to(
						".card-wrapper-3",
						{
							x: 0,
							y: 0,
							rotation: 0,
							scale: 1.1,
							zIndex: 30,
							duration: 1,
						},
						3,
					)
					.to(
						".card-wrapper-4",
						{ x: 120, y: 0, rotation: 5, scale: 1, duration: 1 },
						3,
					)
					.to(
						".card-wrapper-5",
						{ x: 240, y: 0, rotation: 10, scale: 1, duration: 1 },
						3,
					)

					// Content transition
					.to(
						".section-3",
						{ autoAlpha: 0, y: -50, duration: 0.5 },
						3.2,
					)
					.to(
						".section-4",
						{ autoAlpha: 1, y: 0, duration: 0.5 },
						3.5,
					);

				// Entrance Animation
				// Animate wrappers from center with scale
				gsap.fromTo(
					".card-wrapper",
					{
						scale: 0,
						opacity: 0,
					},
					{
						scale: 1,
						opacity: 1,
						duration: 1.2,
						stagger: 0.1,
						ease: "elastic.out(1, 0.7)",
						delay: 0.2,
					},
				);

				// Animate content in
				gsap.fromTo(
					".section-0",
					{
						y: 30,
						autoAlpha: 0,
					},
					{
						y: 0,
						autoAlpha: 1,
						duration: 1,
						delay: 0.8,
						ease: "power3.out",
					},
				);
			}, scrollContainer);

			// Small delay to allow for any final layout shifts/scroll restoration
			setTimeout(() => {
				ScrollTrigger.refresh();
			}, 100);
		};

		init();

		return () => {
			if (ctx) ctx.revert();
			if (ScrollTriggerInstance) {
				ScrollTriggerInstance.getAll().forEach((trigger) => {
					trigger.kill();
				});
			}
		};
	});
</script>

<div class="scroll-container" bind:this={scrollContainer}>
	<section class="hero">
		<div class="hero-background"></div>

		<!-- Memory Cards Container -->
		<div class="cards-container">
			<div class="card-stack">
				{#each memoryCardColors as color, i}
					<!-- Wrapper handles Scroll Position -->
					<div class="card-wrapper card-wrapper-{i + 1}">
						<!-- Inner card handles Float & Hover -->
						<div class="memory-card">
							<canvas
								bind:this={canvasRefs[i]}
								width="180"
								height="240"
							></canvas>
						</div>
					</div>
				{/each}
			</div>
		</div>

		<!-- Content sections that change on scroll -->
		<div class="content-wrapper">
			{#each sections as section, i}
				<div
					class="hero-content section-{i}"
					style="opacity: {i === 0 ? 1 : 0}; visibility: {i === 0
						? 'visible'
						: 'hidden'}"
				>
					<h1 class="hero-title">
						{section.title}
						<span class="gradient-text">{section.highlight}</span>
					</h1>
					<p class="hero-subtitle">
						{section.subtitle}
					</p>
					<div class="hero-buttons">
						<button class="primary-button"
							>{section.buttons[0]}</button
						>
						<button class="secondary-button"
							>{section.buttons[1]}</button
						>
					</div>
				</div>
			{/each}
		</div>
	</section>
</div>

<style>
	.scroll-container {
		height: 400vh; /* Creates scrollable area */
		position: relative;
	}

	.hero {
		position: relative;
		height: 100vh;
		width: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		overflow: hidden;
		background: var(--dark-bg);
		color: var(--text-light);
		perspective: 1000px; /* Adds 3D depth */
	}

	.hero-background {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background: radial-gradient(
				circle at 20% 50%,
				rgba(99, 102, 241, 0.15) 0%,
				transparent 50%
			),
			radial-gradient(
				circle at 80% 80%,
				rgba(139, 92, 246, 0.15) 0%,
				transparent 50%
			);
		z-index: 0;
	}

	.cards-container {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		z-index: 0;
		pointer-events: none;
		width: 100%;
		height: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.card-stack {
		position: relative;
		width: 0;
		height: 0;
		/* Using 0x0 size and letting children position themselves relative to center */
	}

	.card-wrapper {
		position: absolute;
		top: 0;
		left: 0;
		/* Center the wrapper itself so (0,0) is dead center */
		transform: translate(-50%, -50%);
		will-change: transform;
	}

	.memory-card {
		cursor: pointer;
		pointer-events: auto;
		filter: drop-shadow(0 10px 30px rgba(0, 0, 0, 0.5));
		transition: filter 0.3s ease;
		/* Inner card is relative to wrapper */
	}

	.memory-card:hover {
		filter: drop-shadow(0 15px 40px rgba(99, 102, 241, 0.6));
		transform: scale(1.05); /* Simple hover scale */
	}

	.memory-card canvas {
		display: block;
		border-radius: 8px;
	}

	.content-wrapper {
		position: relative;
		z-index: 1;
		width: 100%;
		max-width: 1200px;
		padding: 0 2rem;
		pointer-events: none; /* Let clicks pass through to cards if needed, but buttons need events */
	}

	/* Re-enable pointer events for buttons */
	.hero-buttons {
		pointer-events: auto;
	}

	.hero-content {
		position: absolute;
		left: 50%;
		top: 50%;
		transform: translate(-50%, -50%);
		width: 100%;
		max-width: 600px;
		text-align: center;
		z-index: 10;
		/* opacity and visibility handled by inline styles and GSAP */
	}

	.hero-title {
		font-size: 3.5rem;
		margin-bottom: 1.5rem;
		line-height: 1.2;
		text-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
	}

	.gradient-text {
		background: var(--gradient);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
		display: block;
		padding-bottom: 0.15em;
	}

	.hero-subtitle {
		font-size: 1.25rem;
		color: #cbd5e1;
		margin-bottom: 2.5rem;
		line-height: 1.6;
		text-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
	}

	.hero-buttons {
		display: flex;
		gap: 1rem;
		justify-content: center;
		flex-wrap: wrap;
	}

	.primary-button,
	.secondary-button {
		padding: 1rem 2rem;
		border-radius: 50px;
		font-size: 1rem;
		font-weight: 600;
		transition: all 0.3s ease;
		cursor: pointer;
	}

	.primary-button {
		background: var(--gradient);
		color: white;
		border: none;
	}

	.primary-button:hover {
		transform: translateY(-3px);
		box-shadow: 0 15px 30px rgba(99, 102, 241, 0.4);
	}

	.secondary-button {
		background: rgba(255, 255, 255, 0.05);
		color: white;
		border: 1px solid rgba(255, 255, 255, 0.2);
		backdrop-filter: blur(10px);
	}

	.secondary-button:hover {
		border-color: var(--primary-color);
		background: rgba(99, 102, 241, 0.1);
	}

	@media (max-width: 768px) {
		.hero-title {
			font-size: 2.5rem;
		}

		.hero-subtitle {
			font-size: 1.1rem;
		}

		.memory-card canvas {
			width: 120px;
			height: 160px;
		}

		.content-wrapper {
			padding: 0 1rem;
		}
	}
</style>
