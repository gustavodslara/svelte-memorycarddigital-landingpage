<script>
	import { onMount } from 'svelte';

	const features = [
		{
			icon: '🎨',
			title: 'Custom Design',
			description: 'Create beautiful memory cards with our intuitive design tools and templates.'
		},
		{
			icon: '☁️',
			title: 'Cloud Storage',
			description: 'Safely store all your digital memories in the cloud with unlimited storage.'
		},
		{
			icon: '🔗',
			title: 'Easy Sharing',
			description: 'Share your memory cards with friends and family with a single link.'
		},
		{
			icon: '📱',
			title: 'Mobile Ready',
			description: 'Access and create memory cards on any device, anywhere, anytime.'
		},
		{
			icon: '🎬',
			title: 'Video Support',
			description: 'Add videos and audio to make your memories come alive.'
		},
		{
			icon: '🔒',
			title: 'Privacy First',
			description: 'Your memories are private and secure with end-to-end encryption.'
		}
	];

	// Content sections for scroll
	const sections = [
		{
			title: 'Amazing Features',
			subtitle: 'Everything you need to preserve and share your precious memories',
			featureIndices: [0, 1, 2] // Custom Design, Cloud Storage, Easy Sharing
		},
		{
			title: 'Built For You',
			subtitle: 'Designed with your needs in mind, accessible anywhere',
			featureIndices: [3, 4, 5] // Mobile Ready, Video Support, Privacy First
		}
	];

	onMount(async () => {
		const { default: gsap } = await import('gsap');
		const { ScrollTrigger } = await import('gsap/ScrollTrigger');
		
		gsap.registerPlugin(ScrollTrigger);

		// Pin the features section
		ScrollTrigger.create({
			trigger: '.scroll-container',
			start: 'top top',
			end: '+=200%',
			pin: '.features',
			pinSpacing: false,
			scrub: true
		});

		// Create master timeline for scroll-based animations
		const masterTimeline = gsap.timeline({
			scrollTrigger: {
				trigger: '.scroll-container',
				start: 'top top',
				end: '+=200%',
				scrub: 1
			}
		});

		// Animate feature cards spreading out
		masterTimeline
			.from('.feature-card', {
				scale: 0,
				opacity: 0,
				duration: 0.5,
				stagger: 0.1,
				ease: 'back.out(1.7)'
			}, 0)
			.to('.feature-card-0', { x: -300, y: -150, rotation: -8, duration: 1 }, 0.5)
			.to('.feature-card-1', { x: 0, y: -180, rotation: 0, duration: 1 }, 0.5)
			.to('.feature-card-2', { x: 300, y: -150, rotation: 8, duration: 1 }, 0.5)
			.to('.feature-card-3', { x: -300, y: 150, rotation: -5, duration: 1 }, 0.5)
			.to('.feature-card-4', { x: 0, y: 180, rotation: 0, duration: 1 }, 0.5)
			.to('.feature-card-5', { x: 300, y: 150, rotation: 5, duration: 1 }, 0.5)
			.to('.section-0', { opacity: 0, y: -50, duration: 0.5 }, 0.8)
			.to('.section-1', { opacity: 1, y: 0, duration: 0.5 }, 1);

		// Continuous floating animation for feature cards
		document.querySelectorAll('.feature-card').forEach((card, index) => {
			gsap.to(card, {
				y: `+=${gsap.utils.random(-20, 20)}`,
				rotation: `+=${gsap.utils.random(-5, 5)}`,
				duration: gsap.utils.random(3, 5),
				repeat: -1,
				yoyo: true,
				ease: 'sine.inOut',
				delay: index * 0.2
			});
		});

		// Initial entrance
		gsap.from('.section-0', {
			opacity: 0,
			y: 50,
			duration: 1,
			delay: 0.3,
			ease: 'power4.out'
		});
	});
</script>

<div class="scroll-container">
	<section id="features" class="features">
		<div class="features-background"></div>
		
		<!-- Feature Cards - Fixed position -->
		<div class="cards-container">
			<div class="features-grid">
				{#each features as feature, i}
					<div class="feature-card feature-card-{i}">
						<div class="feature-icon">{feature.icon}</div>
						<h3 class="feature-title">{feature.title}</h3>
						<p class="feature-description">{feature.description}</p>
					</div>
				{/each}
			</div>
		</div>

		<!-- Content sections that change on scroll -->
		<div class="content-wrapper">
			{#each sections as section, i}
				<div class="section-content section-{i}" class:active={i === 0}>
					<h2 class="section-title">{section.title}</h2>
					<p class="section-subtitle">{section.subtitle}</p>
				</div>
			{/each}
		</div>
	</section>
</div>

<style>
	.scroll-container {
		height: 300vh;
		position: relative;
	}

	.features {
		position: relative;
		height: 100vh;
		width: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		overflow: hidden;
		background: var(--light-bg);
	}

	.features-background {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background: 
			radial-gradient(circle at 80% 20%, rgba(99, 102, 241, 0.08) 0%, transparent 50%),
			radial-gradient(circle at 20% 80%, rgba(236, 72, 153, 0.08) 0%, transparent 50%);
		z-index: 0;
	}

	.cards-container {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		z-index: 2;
		width: 100%;
		max-width: 1400px;
		padding: 0 2rem;
	}

	.features-grid {
		position: relative;
		display: flex;
		align-items: center;
		justify-content: center;
		min-height: 400px;
	}

	.feature-card {
		position: absolute;
		background: white;
		padding: 2rem;
		border-radius: 20px;
		box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
		border: 1px solid rgba(99, 102, 241, 0.1);
		width: 280px;
		transition: all 0.3s ease;
		cursor: pointer;
		transform-origin: center center;
	}

	.feature-card:hover {
		box-shadow: 0 20px 40px rgba(99, 102, 241, 0.25);
		border-color: var(--primary-color);
		z-index: 10;
	}

	.feature-icon {
		font-size: 3rem;
		margin-bottom: 1rem;
	}

	.feature-title {
		font-size: 1.4rem;
		margin-bottom: 0.75rem;
		color: var(--text-dark);
	}

	.feature-description {
		color: #64748b;
		line-height: 1.6;
		font-size: 0.95rem;
	}

	.content-wrapper {
		position: relative;
		z-index: 1;
		width: 100%;
		text-align: center;
		padding: 0 2rem;
	}

	.section-content {
		position: absolute;
		left: 50%;
		top: 10%;
		transform: translateX(-50%);
		width: 100%;
		max-width: 800px;
		opacity: 0;
	}

	.section-content.active,
	.section-content.section-0 {
		opacity: 1;
	}

	.section-title {
		font-size: 3rem;
		margin-bottom: 1rem;
		color: var(--text-dark);
		background: var(--gradient);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}

	.section-subtitle {
		font-size: 1.25rem;
		color: #64748b;
	}

	@media (max-width: 768px) {
		.section-title {
			font-size: 2rem;
		}

		.section-subtitle {
			font-size: 1rem;
		}

		.feature-card {
			width: 200px;
			padding: 1.5rem;
		}

		.feature-icon {
			font-size: 2rem;
		}

		.feature-title {
			font-size: 1.1rem;
		}

		.feature-description {
			font-size: 0.85rem;
		}
	}
</style>
