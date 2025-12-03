<script>
	import { onMount } from "svelte";
	import { t } from "svelte-i18n";

	let benefits = $derived([
		$t("about.benefits.backup"),
		$t("about.benefits.sync"),
		$t("about.benefits.cloud"),
		$t("about.benefits.support"),
	]);

	let sections = $derived([
		{
			title: $t("about.easyBackup.title"),
			highlight: $t("about.easyBackup.highlight"),
			subtitle: $t("about.easyBackup.subtitle"),
			description: $t("about.easyBackup.description"),
		},
		{
			title: $t("about.howItWorks.title"),
			highlight: $t("about.howItWorks.highlight"),
			subtitle: $t("about.howItWorks.subtitle"),
			description: $t("about.howItWorks.description"),
			steps: [
				$t("about.steps.step1"),
				$t("about.steps.step2"),
				$t("about.steps.step3"),
				$t("about.steps.step4"),
				$t("about.steps.step5"),
			],
		},
	]);

	// PS2 Memory Card colors (subset for About section)
	const memoryCardColors = [
		{ primary: "#1e40af", secondary: "#1e3a8a", accent: "#3b82f6" }, // Blue
		{ primary: "#7c2d12", secondary: "#991b1b", accent: "#dc2626" }, // Red
		{ primary: "#065f46", secondary: "#064e3b", accent: "#10b981" }, // Green
	];

	/** @type {HTMLCanvasElement[]} */
	let canvasRefs = [];

	/**
	 * @param {HTMLCanvasElement} canvas
	 * @param {{ primary: string; secondary: string; accent: string; }} colorScheme
	 */
	function drawMemoryCard(canvas, colorScheme) {
		const ctx = canvas.getContext("2d");
		if (!ctx) return;
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
		(async () => {
			const { default: gsap } = await import("gsap");
			const { ScrollTrigger } = await import("gsap/ScrollTrigger");

			gsap.registerPlugin(ScrollTrigger);

			// Draw memory cards
			canvasRefs.forEach((canvas, index) => {
				if (canvas) {
					drawMemoryCard(canvas, memoryCardColors[index]);
				}
			});

			ctx = gsap.context(() => {
				// Pin the about section with improved configuration
				ScrollTrigger.create({
					trigger: ".scroll-container",
					start: "top top",
					end: "+=200%",
					pin: ".about",
					pinSpacing: false,
					scrub: true,
					anticipatePin: 1, // Prevents layout shifts
					invalidateOnRefresh: true, // Recalculates on resize/refresh
				});

				// Continuous floating animation for elements (inner)
				gsap.to(".floating-element-1", {
					x: 40,
					y: -30,
					rotation: 15,
					duration: 4,
					repeat: -1,
					yoyo: true,
					ease: "sine.inOut",
				});

				gsap.to(".floating-element-2", {
					x: -30,
					y: 40,
					rotation: -12,
					duration: 5,
					repeat: -1,
					yoyo: true,
					ease: "sine.inOut",
					delay: 0.5,
				});

				gsap.to(".floating-element-3", {
					x: 25,
					y: 35,
					rotation: 8,
					duration: 4.5,
					repeat: -1,
					yoyo: true,
					ease: "sine.inOut",
					delay: 1,
				});

				// Create master timeline for scroll-based animations
				const masterTimeline = gsap.timeline({
					scrollTrigger: {
						trigger: ".scroll-container",
						start: "top top",
						end: "+=200%",
						scrub: 1,
						invalidateOnRefresh: true,
					},
				});

				// Animate floating elements spreading (outer wrappers)
				// Set initial state explicitly to prevent glitches
				gsap.set(
					".floating-wrapper-1, .floating-wrapper-2, .floating-wrapper-3",
					{
						x: 0,
						y: 0,
						scale: 1,
					},
				);

				gsap.set(".section-0", { opacity: 1, y: 0 });
				gsap.set(".section-1", { opacity: 0, y: 50 });

				masterTimeline
					.to(
						".floating-wrapper-1",
						{
							x: -200,
							y: -150,
							scale: 1.2,
							duration: 1,
							ease: "power2.inOut",
						},
						0,
					)
					.to(
						".floating-wrapper-2",
						{
							x: 250,
							y: -100,
							scale: 0.9,
							duration: 1,
							ease: "power2.inOut",
						},
						0,
					)
					.to(
						".floating-wrapper-3",
						{
							x: -150,
							y: 200,
							scale: 1.1,
							duration: 1,
							ease: "power2.inOut",
						},
						0,
					)
					.to(
						".section-0",
						{
							opacity: 0,
							y: -50,
							duration: 0.5,
							ease: "power2.inOut",
						},
						0.8,
					)
					.to(
						".section-1",
						{
							opacity: 1,
							y: 0,
							duration: 0.5,
							ease: "power2.inOut",
						},
						1,
					);

				// Initial entrance
				gsap.from(".section-0", {
					opacity: 0,
					y: 50,
					duration: 1,
					delay: 0.3,
					ease: "power4.out",
				});

				gsap.from(
					".floating-wrapper-1, .floating-wrapper-2, .floating-wrapper-3",
					{
						scale: 0,
						opacity: 0,
						duration: 1,
						delay: 0.5,
						stagger: 0.2,
						ease: "back.out(1.7)",
					},
				);

				gsap.from(".benefit-item", {
					x: -50,
					opacity: 0,
					duration: 0.6,
					delay: 1,
					stagger: 0.15,
					ease: "power3.out",
				});
			});
		})();

		return () => {
			// Proper cleanup
			if (ctx) {
				ctx.revert();
			}
			// Kill all ScrollTriggers associated with this component
			ScrollTrigger.getAll().forEach((trigger) => {
				if (trigger.vars.trigger === ".scroll-container") {
					trigger.kill();
				}
			});
		};
	});
</script>

<div class="scroll-container">
	<section id="about" class="about">
		<div class="about-background"></div>

		<!-- Floating visual elements -->
		<div class="visuals-container">
			<div class="floating-wrapper floating-wrapper-1">
				<div class="floating-element floating-element-1">
					<canvas bind:this={canvasRefs[0]} width="180" height="240"
					></canvas>
				</div>
			</div>
			<div class="floating-wrapper floating-wrapper-2">
				<div class="floating-element floating-element-2">
					<canvas bind:this={canvasRefs[1]} width="180" height="240"
					></canvas>
				</div>
			</div>
			<div class="floating-wrapper floating-wrapper-3">
				<div class="floating-element floating-element-3">
					<canvas bind:this={canvasRefs[2]} width="180" height="240"
					></canvas>
				</div>
			</div>
		</div>

		<!-- Content sections that change on scroll -->
		<div class="content-wrapper">
			{#each sections as section, i}
				<div class="about-content section-{i}" class:active={i === 0}>
					<h2 class="about-title">
						{section.title}
						<span class="gradient-text">{section.highlight}</span>
					</h2>
					<p class="about-subtitle">{section.subtitle}</p>
					<p class="about-text">{section.description}</p>

					{#if i === 0}
						<ul class="about-list">
							{#each benefits as benefit}
								<li class="benefit-item">{benefit}</li>
							{/each}
						</ul>
						<a href="/download" class="learn-more-button"
							>{$t("about.downloadNow")}</a
						>
					{/if}
					{#if i === 1 && section.steps}
						<ol class="steps-list">
							{#each section.steps as step}
								<li class="step-item">{step}</li>
							{/each}
						</ol>
						<a href="/download" class="learn-more-button"
							>{$t("about.downloadNow")}</a
						>
					{/if}
				</div>
			{/each}
		</div>
	</section>
</div>

<style>
	.scroll-container {
		height: 300vh;
		position: relative;
		background: var(--dark-bg); /* Prevent white flash */
	}

	.about {
		position: relative;
		height: 100vh;
		width: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		overflow: hidden;
		background: var(--dark-bg);
		color: var(--text-light);
	}

	.about-background {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background: radial-gradient(
				circle at 30% 40%,
				rgba(99, 102, 241, 0.15) 0%,
				transparent 50%
			),
			radial-gradient(
				circle at 70% 60%,
				rgba(236, 72, 153, 0.15) 0%,
				transparent 50%
			);
		z-index: 0;
	}

	.visuals-container {
		position: absolute;
		top: 50%;
		right: 10%;
		transform: translateY(-50%);
		z-index: 1;
		width: 400px;
		height: 400px;
	}

	.floating-wrapper {
		position: absolute;
		will-change: transform, opacity; /* Optimize animations */
	}

	.floating-element {
		width: 100%;
		height: 100%;
		/* border-radius: 20px; Removed as canvas handles shape */
		/* filter: blur(1px); Removed blur for sharper cards */
		transform-origin: center center;
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.floating-element canvas {
		filter: drop-shadow(0 20px 40px rgba(0, 0, 0, 0.3));
		border-radius: 8px;
	}

	.floating-wrapper-2 {
		width: 150px;
		height: 150px;
		top: 100px;
		right: 50px;
	}

	.floating-wrapper-3 {
		width: 180px;
		height: 180px;
		bottom: 50px;
		left: 100px;
	}

	.content-wrapper {
		position: relative;
		z-index: 2;
		width: 100%;
		max-width: 1200px;
		padding: 0 2rem;
	}

	.about-content {
		position: absolute;
		left: 5%;
		top: 50%;
		transform: translateY(-50%);
		max-width: 600px;
		opacity: 0;
		will-change: opacity, transform; /* Optimize animations */
	}

	.about-content.active,
	.about-content.section-0 {
		opacity: 1;
	}

	.about-title {
		font-size: 2.5rem;
		margin-bottom: 1.5rem;
		line-height: 1.2;
	}

	.gradient-text {
		background: var(--gradient);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
		display: block;
	}

	.about-subtitle {
		font-size: 1.3rem;
		color: #e2e8f0;
		margin-bottom: 1rem;
		font-weight: 500;
	}

	.about-text {
		font-size: 1.1rem;
		color: #cbd5e1;
		margin-bottom: 1.5rem;
		line-height: 1.8;
	}

	.about-list {
		list-style: none;
		margin: 2rem 0;
	}

	.benefit-item {
		font-size: 1.1rem;
		padding: 0.5rem 0;
		color: #cbd5e1;
	}

	.steps-list {
		list-style: decimal;
		margin: 2rem 0;
		padding-left: 1.5rem;
	}

	.step-item {
		font-size: 1.1rem;
		padding: 0.75rem 0;
		color: #cbd5e1;
		line-height: 1.6;
	}

	.learn-more-button {
		display: inline-block;
		background: var(--gradient);
		color: white;
		padding: 1rem 2.5rem;
		border-radius: 50px;
		font-size: 1rem;
		font-weight: 600;
		margin-top: 1rem;
		border: none;
		cursor: pointer;
		transition: all 0.3s ease;
		text-decoration: none;
	}

	.learn-more-button:hover {
		transform: translateY(-3px);
		box-shadow: 0 15px 30px rgba(99, 102, 241, 0.4);
	}

	@media (max-width: 768px) {
		.about-title {
			font-size: 2rem;
		}

		.about-subtitle {
			font-size: 1.1rem;
		}

		.about-text {
			font-size: 1rem;
		}

		.visuals-container {
			right: 50%;
			transform: translate(50%, -50%);
			opacity: 0.3;
			width: 300px;
			height: 300px;
		}

		.floating-wrapper-1,
		.floating-wrapper-2,
		.floating-wrapper-3 {
			width: 100px;
			height: 100px;
		}

		.about-content {
			left: 50%;
			transform: translate(-50%, -50%);
			text-align: center;
		}
	}
</style>
