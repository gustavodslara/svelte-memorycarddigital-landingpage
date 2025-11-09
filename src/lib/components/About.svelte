<script>
	import { onMount } from 'svelte';

	const benefits = [
		'✨ Beautiful templates and customization options',
		'🚀 Lightning-fast performance',
		'💝 Share with unlimited recipients',
		'🌍 Access from anywhere in the world'
	];

	const sections = [
		{
			title: 'Why Choose',
			highlight: 'Memory Card Digital?',
			subtitle: 'We believe that every memory deserves to be preserved in the most beautiful way possible.',
			description: 'Our platform combines cutting-edge technology with intuitive design to help you create stunning digital memory cards that will last forever.'
		},
		{
			title: 'Moments That',
			highlight: 'Matter',
			subtitle: 'Whether it\'s a special occasion or everyday moments...',
			description: 'Memory Card Digital makes it easy to capture, customize, and share your memories with the people who matter most.'
		}
	];

	onMount(async () => {
		const { default: gsap } = await import('gsap');
		const { ScrollTrigger } = await import('gsap/ScrollTrigger');
		
		gsap.registerPlugin(ScrollTrigger);

		// Pin the about section
		ScrollTrigger.create({
			trigger: '.scroll-container',
			start: 'top top',
			end: '+=200%',
			pin: '.about',
			pinSpacing: false,
			scrub: true
		});

		// Continuous floating animation for elements
		gsap.to('.floating-element-1', {
			x: 40,
			y: -30,
			rotation: 15,
			duration: 4,
			repeat: -1,
			yoyo: true,
			ease: 'sine.inOut'
		});

		gsap.to('.floating-element-2', {
			x: -30,
			y: 40,
			rotation: -12,
			duration: 5,
			repeat: -1,
			yoyo: true,
			ease: 'sine.inOut',
			delay: 0.5
		});

		gsap.to('.floating-element-3', {
			x: 25,
			y: 35,
			rotation: 8,
			duration: 4.5,
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

		// Animate floating elements spreading
		masterTimeline
			.to('.floating-element-1', { x: -200, y: -150, scale: 1.2, duration: 1 }, 0)
			.to('.floating-element-2', { x: 250, y: -100, scale: 0.9, duration: 1 }, 0)
			.to('.floating-element-3', { x: -150, y: 200, scale: 1.1, duration: 1 }, 0)
			.to('.section-0', { opacity: 0, y: -50, duration: 0.5 }, 0.8)
			.to('.section-1', { opacity: 1, y: 0, duration: 0.5 }, 1);

		// Initial entrance
		gsap.from('.section-0', {
			opacity: 0,
			y: 50,
			duration: 1,
			delay: 0.3,
			ease: 'power4.out'
		});

		gsap.from('.floating-element', {
			scale: 0,
			opacity: 0,
			duration: 1,
			delay: 0.5,
			stagger: 0.2,
			ease: 'back.out(1.7)'
		});

		gsap.from('.benefit-item', {
			x: -50,
			opacity: 0,
			duration: 0.6,
			delay: 1,
			stagger: 0.15,
			ease: 'power3.out'
		});
	});
</script>

<div class="scroll-container">
	<section id="about" class="about">
		<div class="about-background"></div>
		
		<!-- Floating visual elements -->
		<div class="visuals-container">
			<div class="floating-element floating-element-1"></div>
			<div class="floating-element floating-element-2"></div>
			<div class="floating-element floating-element-3"></div>
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
						<button class="learn-more-button">Learn More</button>
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
		background: 
			radial-gradient(circle at 30% 40%, rgba(99, 102, 241, 0.15) 0%, transparent 50%),
			radial-gradient(circle at 70% 60%, rgba(236, 72, 153, 0.15) 0%, transparent 50%);
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

	.floating-element {
		position: absolute;
		border-radius: 20px;
		filter: blur(1px);
		transform-origin: center center;
	}

	.floating-element-1 {
		width: 200px;
		height: 200px;
		background: linear-gradient(135deg, rgba(99, 102, 241, 0.6), rgba(139, 92, 246, 0.6));
		top: 50px;
		left: 50px;
		box-shadow: 0 20px 60px rgba(99, 102, 241, 0.4);
	}

	.floating-element-2 {
		width: 150px;
		height: 150px;
		background: linear-gradient(135deg, rgba(236, 72, 153, 0.6), rgba(139, 92, 246, 0.6));
		top: 100px;
		right: 50px;
		box-shadow: 0 20px 60px rgba(236, 72, 153, 0.4);
	}

	.floating-element-3 {
		width: 180px;
		height: 180px;
		background: linear-gradient(135deg, rgba(99, 102, 241, 0.6), rgba(236, 72, 153, 0.6));
		bottom: 50px;
		left: 100px;
		box-shadow: 0 20px 60px rgba(139, 92, 246, 0.4);
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

	.learn-more-button {
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

		.floating-element-1,
		.floating-element-2,
		.floating-element-3 {
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
