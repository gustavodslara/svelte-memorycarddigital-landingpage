<script>
	import { onMount } from 'svelte';

	let canvasRefs = [];
	let scrollContainer;

	// PS2 Memory Card colors (minimalist solid colors)
	const memoryCardColors = [
		{ primary: '#1e40af', secondary: '#1e3a8a', accent: '#3b82f6' }, // Blue
		{ primary: '#7c2d12', secondary: '#991b1b', accent: '#dc2626' }, // Red
		{ primary: '#065f46', secondary: '#064e3b', accent: '#10b981' }, // Green
		{ primary: '#1f2937', secondary: '#111827', accent: '#4b5563' }  // Dark Gray
	];

	// Content for different scroll sections
	const sections = [
		{
			title: 'Transform Your Digital',
			highlight: 'Memories',
			subtitle: 'Create stunning digital memory cards that preserve your precious moments forever. Share, customize, and relive your memories in a whole new way.',
			buttons: ['Start Free Trial', 'Watch Demo']
		},
		{
			title: 'Store Every',
			highlight: 'Moment',
			subtitle: 'Unlimited storage for all your precious memories. Never lose a moment with our secure cloud backup and easy organization.',
			buttons: ['Learn More', 'Get Started']
		},
		{
			title: 'Share With',
			highlight: 'Everyone',
			subtitle: 'Share your memory cards with friends and family. Create collaborative albums and relive moments together.',
			buttons: ['Try Sharing', 'View Features']
		},
		{
			title: 'Preserve',
			highlight: 'Forever',
			subtitle: 'Your memories, secured for generations. Advanced encryption and backup ensure your moments last forever.',
			buttons: ['Start Now', 'Contact Us']
		}
	];

	function drawMemoryCard(canvas, colorScheme) {
		const ctx = canvas.getContext('2d');
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

	onMount(async () => {
		const { default: gsap } = await import('gsap');
		const { ScrollTrigger } = await import('gsap/ScrollTrigger');
		
		gsap.registerPlugin(ScrollTrigger);

		// Draw memory cards on canvas
		canvasRefs.forEach((canvas, index) => {
			if (canvas) {
				drawMemoryCard(canvas, memoryCardColors[index]);
			}
		});

		// Continuous floating animation for each card (never stops)
		const cards = document.querySelectorAll('.memory-card');
		cards.forEach((card, index) => {
			gsap.to(card, {
				x: `+=${gsap.utils.random(-30, 30)}`,
				y: `+=${gsap.utils.random(-30, 30)}`,
				rotation: `+=${gsap.utils.random(-10, 10)}`,
				duration: gsap.utils.random(3, 5),
				repeat: -1,
				yoyo: true,
				ease: 'sine.inOut',
				delay: index * 0.3
			});
		});

		// Pin the entire hero section
		ScrollTrigger.create({
			trigger: '.scroll-container',
			start: 'top top',
			end: '+=300%',
			pin: '.hero',
			pinSpacing: false,
			scrub: true
		});

		// Create master timeline for scroll-based animations
		const masterTimeline = gsap.timeline({
			scrollTrigger: {
				trigger: '.scroll-container',
				start: 'top top',
				end: '+=300%',
				scrub: 1
			}
		});

		// Section 0 to 1: Cards start spreading
		masterTimeline
			.to('.memory-card-1', { x: -150, y: -100, rotation: -25, duration: 1 }, 0)
			.to('.memory-card-2', { x: 150, y: -120, rotation: 20, duration: 1 }, 0)
			.to('.memory-card-3', { x: -120, y: 130, rotation: -18, duration: 1 }, 0)
			.to('.memory-card-4', { x: 180, y: 110, rotation: 28, duration: 1 }, 0)
			.to('.section-0', { opacity: 0, y: -50, duration: 0.5 }, 0.3)
			.to('.section-1', { opacity: 1, y: 0, duration: 0.5 }, 0.5);

		// Section 1 to 2: Cards spread more
		masterTimeline
			.to('.memory-card-1', { x: -280, y: -200, rotation: -40, duration: 1 }, 1)
			.to('.memory-card-2', { x: 300, y: -180, rotation: 35, duration: 1 }, 1)
			.to('.memory-card-3', { x: -250, y: 240, rotation: -32, duration: 1 }, 1)
			.to('.memory-card-4', { x: 320, y: 220, rotation: 42, duration: 1 }, 1)
			.to('.section-1', { opacity: 0, y: -50, duration: 0.5 }, 1.3)
			.to('.section-2', { opacity: 1, y: 0, duration: 0.5 }, 1.5);

		// Section 2 to 3: Cards at maximum spread
		masterTimeline
			.to('.memory-card-1', { x: -400, y: -280, rotation: -55, scale: 0.9, duration: 1 }, 2)
			.to('.memory-card-2', { x: 420, y: -250, rotation: 48, scale: 0.95, duration: 1 }, 2)
			.to('.memory-card-3', { x: -380, y: 340, rotation: -45, scale: 0.85, duration: 1 }, 2)
			.to('.memory-card-4', { x: 450, y: 320, rotation: 58, scale: 1, duration: 1 }, 2)
			.to('.section-2', { opacity: 0, y: -50, duration: 0.5 }, 2.3)
			.to('.section-3', { opacity: 1, y: 0, duration: 0.5 }, 2.5);

		// Initial entrance animation
		gsap.from('.section-0', {
			opacity: 0,
			y: 50,
			duration: 1,
			delay: 0.5,
			ease: 'power4.out'
		});

		gsap.from('.memory-card', {
			scale: 0,
			opacity: 0,
			duration: 1,
			delay: 0.8,
			stagger: 0.15,
			ease: 'back.out(1.7)'
		});
	});
</script>

<div class="scroll-container" bind:this={scrollContainer}>
	<section class="hero">
		<div class="hero-background"></div>
		
		<!-- Memory Cards - Fixed position -->
		<div class="cards-container">
			<div class="card-stack">
				{#each memoryCardColors as color, i}
					<div class="memory-card memory-card-{i + 1}">
						<canvas 
							bind:this={canvasRefs[i]} 
							width="180" 
							height="240"
						></canvas>
					</div>
				{/each}
			</div>
		</div>

		<!-- Content sections that change on scroll -->
		<div class="content-wrapper">
			{#each sections as section, i}
				<div class="hero-content section-{i}" class:active={i === 0}>
					<h1 class="hero-title">
						{section.title}
						<span class="gradient-text">{section.highlight}</span>
					</h1>
					<p class="hero-subtitle">
						{section.subtitle}
					</p>
					<div class="hero-buttons">
						<button class="primary-button">{section.buttons[0]}</button>
						<button class="secondary-button">{section.buttons[1]}</button>
					</div>
				</div>
			{/each}
		</div>
	</section>
</div>

<style>
	.scroll-container {
		height: 300vh; /* Creates scrollable area */
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
	}

	.hero-background {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background: 
			radial-gradient(circle at 20% 50%, rgba(99, 102, 241, 0.15) 0%, transparent 50%),
			radial-gradient(circle at 80% 80%, rgba(139, 92, 246, 0.15) 0%, transparent 50%);
		z-index: 0;
	}

	.cards-container {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		z-index: 2;
		pointer-events: none;
	}

	.card-stack {
		position: relative;
		width: 400px;
		height: 400px;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.memory-card {
		position: absolute;
		cursor: pointer;
		pointer-events: auto;
		filter: drop-shadow(0 10px 30px rgba(0, 0, 0, 0.5));
		transition: filter 0.3s ease;
	}

	.memory-card:hover {
		filter: drop-shadow(0 15px 40px rgba(99, 102, 241, 0.6));
	}

	.memory-card canvas {
		display: block;
		border-radius: 8px;
	}

	.memory-card-1,
	.memory-card-2,
	.memory-card-3,
	.memory-card-4 {
		transform-origin: center center;
	}

	.content-wrapper {
		position: relative;
		z-index: 1;
		width: 100%;
		max-width: 1200px;
		padding: 0 2rem;
	}

	.hero-content {
		position: absolute;
		left: 50%;
		top: 50%;
		transform: translate(-50%, -50%);
		width: 100%;
		max-width: 600px;
		text-align: center;
		opacity: 0;
		transition: none;
	}

	.hero-content.active,
	.hero-content.section-0 {
		opacity: 1;
	}

	.hero-title {
		font-size: 3.5rem;
		margin-bottom: 1.5rem;
		line-height: 1.1;
	}

	.gradient-text {
		background: var(--gradient);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
		display: block;
	}

	.hero-subtitle {
		font-size: 1.25rem;
		color: #cbd5e1;
		margin-bottom: 2.5rem;
		line-height: 1.6;
	}

	.hero-buttons {
		display: flex;
		gap: 1rem;
		justify-content: center;
		flex-wrap: wrap;
	}

	.primary-button, .secondary-button {
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
		background: transparent;
		color: white;
		border: 2px solid rgba(255, 255, 255, 0.3);
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

		.card-stack {
			width: 280px;
			height: 280px;
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
