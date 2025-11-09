<script>
	import { onMount } from 'svelte';

	const sections = [
		{
			title: 'Ready to Preserve',
			highlight: 'Your Memories?',
			subtitle: 'Join thousands of users who trust Memory Card Digital to keep their precious moments safe.',
			buttons: ['Start Free Trial', 'Contact Sales']
		},
		{
			title: 'Start Your',
			highlight: 'Journey Today',
			subtitle: 'Create unlimited memory cards, share with everyone, and preserve your moments forever.',
			buttons: ['Get Started Now', 'View Pricing']
		}
	];

	onMount(async () => {
		const { default: gsap } = await import('gsap');
		const { ScrollTrigger } = await import('gsap/ScrollTrigger');
		
		gsap.registerPlugin(ScrollTrigger);

		// Pin the CTA section
		ScrollTrigger.create({
			trigger: '.scroll-container',
			start: 'top top',
			end: '+=200%',
			pin: '.cta',
			pinSpacing: false,
			scrub: true
		});

		// Continuous rotation animation for decorative circles
		gsap.to('.circle-1', {
			rotation: 360,
			duration: 20,
			repeat: -1,
			ease: 'none'
		});

		gsap.to('.circle-2', {
			rotation: -360,
			duration: 25,
			repeat: -1,
			ease: 'none'
		});

		// Floating animation
		gsap.to('.circle-1', {
			y: -30,
			x: 20,
			duration: 4,
			repeat: -1,
			yoyo: true,
			ease: 'sine.inOut'
		});

		gsap.to('.circle-2', {
			y: 40,
			x: -30,
			duration: 5,
			repeat: -1,
			yoyo: true,
			ease: 'sine.inOut',
			delay: 1
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

		// Animate circles moving and scaling
		masterTimeline
			.to('.circle-1', { x: -300, y: -200, scale: 1.5, opacity: 0.3, duration: 1 }, 0)
			.to('.circle-2', { x: 350, y: 250, scale: 1.3, opacity: 0.4, duration: 1 }, 0)
			.to('.section-0', { opacity: 0, y: -50, duration: 0.5 }, 0.8)
			.to('.section-1', { opacity: 1, y: 0, duration: 0.5 }, 1);

		// Initial entrance
		gsap.from('.section-0', {
			opacity: 0,
			scale: 0.8,
			duration: 1,
			delay: 0.3,
			ease: 'back.out(1.7)'
		});

		gsap.from('.circle-1, .circle-2', {
			scale: 0,
			opacity: 0,
			duration: 1.2,
			delay: 0.5,
			stagger: 0.2,
			ease: 'back.out(1.7)'
		});
	});
</script>

<div class="scroll-container">
	<section id="contact" class="cta">
		<div class="cta-background">
			<div class="circle circle-1"></div>
			<div class="circle circle-2"></div>
		</div>

		<!-- Content sections that change on scroll -->
		<div class="content-wrapper">
			{#each sections as section, i}
				<div class="cta-content section-{i}" class:active={i === 0}>
					<h2 class="cta-title">
						{section.title}
						<span class="highlight-text">{section.highlight}</span>
					</h2>
					<p class="cta-subtitle">{section.subtitle}</p>
					<div class="cta-buttons">
						<button class="cta-primary">{section.buttons[0]}</button>
						<button class="cta-secondary">{section.buttons[1]}</button>
					</div>
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

	.cta {
		position: relative;
		height: 100vh;
		width: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		overflow: hidden;
		background: linear-gradient(135deg, var(--primary-color), var(--secondary-color));
	}

	.cta-background {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		overflow: hidden;
	}

	.circle {
		position: absolute;
		border-radius: 50%;
		background: radial-gradient(circle, rgba(255, 255, 255, 0.2) 0%, rgba(255, 255, 255, 0.05) 70%);
		filter: blur(2px);
	}

	.circle-1 {
		width: 500px;
		height: 500px;
		top: -10%;
		right: -5%;
	}

	.circle-2 {
		width: 600px;
		height: 600px;
		bottom: -15%;
		left: -10%;
	}

	.content-wrapper {
		position: relative;
		z-index: 2;
		width: 100%;
		max-width: 800px;
		padding: 0 2rem;
	}

	.cta-content {
		position: absolute;
		left: 50%;
		top: 50%;
		transform: translate(-50%, -50%);
		width: 100%;
		text-align: center;
		opacity: 0;
	}

	.cta-content.active,
	.cta-content.section-0 {
		opacity: 1;
	}

	.cta-title {
		font-size: 3rem;
		color: white;
		margin-bottom: 1.5rem;
		line-height: 1.2;
	}

	.highlight-text {
		display: block;
		font-weight: 700;
		text-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
	}

	.cta-subtitle {
		font-size: 1.25rem;
		color: rgba(255, 255, 255, 0.95);
		margin-bottom: 3rem;
		line-height: 1.6;
	}

	.cta-buttons {
		display: flex;
		gap: 1rem;
		justify-content: center;
		flex-wrap: wrap;
	}

	.cta-primary, .cta-secondary {
		padding: 1.25rem 2.5rem;
		border-radius: 50px;
		font-size: 1.1rem;
		font-weight: 600;
		cursor: pointer;
		transition: all 0.3s ease;
	}

	.cta-primary {
		background: white;
		color: var(--primary-color);
		border: none;
	}

	.cta-primary:hover {
		transform: translateY(-3px);
		box-shadow: 0 15px 30px rgba(0, 0, 0, 0.3);
	}

	.cta-secondary {
		background: transparent;
		color: white;
		border: 2px solid white;
	}

	.cta-secondary:hover {
		background: rgba(255, 255, 255, 0.15);
		transform: translateY(-3px);
	}

	@media (max-width: 768px) {
		.cta-title {
			font-size: 2rem;
		}

		.cta-subtitle {
			font-size: 1rem;
		}

		.cta-buttons {
			flex-direction: column;
			align-items: center;
		}

		.cta-primary, .cta-secondary {
			width: 100%;
			max-width: 300px;
		}

		.circle-1, .circle-2 {
			width: 300px;
			height: 300px;
		}
	}
</style>
