<script>
	import { onMount } from "svelte";
	import gsap from "gsap";
	import { t } from "svelte-i18n";

	onMount(() => {
		// Animate squares
		gsap.to(".square-1", {
			rotation: 360,
			duration: 20,
			repeat: -1,
			ease: "linear",
		});

		gsap.to(".square-2", {
			rotation: -360,
			duration: 25,
			repeat: -1,
			ease: "linear",
		});

		// Floating animation
		gsap.to(".square-wrapper", {
			y: 20,
			duration: 2,
			yoyo: true,
			repeat: -1,
			ease: "sine.inOut",
			stagger: 0.5,
		});
	});
</script>

<section id="download" class="cta">
	<div class="cta-background">
		<div class="square-wrapper square-wrapper-1">
			<div class="square square-1"></div>
		</div>
		<div class="square-wrapper square-wrapper-2">
			<div class="square square-2"></div>
		</div>
	</div>

	<div class="content-wrapper">
		<div class="cta-content">
			<h2 class="cta-title">
				{$t("cta.title.part1")}
				<span class="highlight-text">{$t("cta.title.highlight")}</span>
			</h2>
			<p class="cta-subtitle">
				{$t("cta.subtitle")}
			</p>
			<div class="cta-buttons">
				<button class="cta-primary">{$t("cta.buttonPrimary")}</button>
				<button class="cta-secondary"
					>{$t("cta.buttonSecondary")}</button
				>
			</div>
		</div>
	</div>
</section>

<style>
	.cta {
		position: relative;
		height: 100vh;
		width: 100%;
		display: flex;
		align-items: center;
		justify-content: center;
		overflow: hidden;
		background: linear-gradient(
			135deg,
			var(--primary-color),
			var(--secondary-color)
		);
	}

	.cta-background {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		overflow: hidden;
	}

	.square-wrapper {
		position: absolute;
	}

	.square {
		width: 100%;
		height: 100%;
		background: linear-gradient(
			135deg,
			rgba(255, 255, 255, 0.2) 0%,
			rgba(255, 255, 255, 0.05) 100%
		);
		backdrop-filter: blur(5px);
		border: 1px solid rgba(255, 255, 255, 0.1);
		border-radius: 20px;
	}

	.square-wrapper-1 {
		width: 400px;
		height: 400px;
		top: 10%;
		right: 5%;
		opacity: 0.6;
	}

	.square-wrapper-2 {
		width: 500px;
		height: 500px;
		bottom: 5%;
		left: -5%;
		opacity: 0.4;
	}

	.content-wrapper {
		position: relative;
		z-index: 2;
		width: 100%;
		max-width: 800px;
		padding: 0 2rem;
	}

	.cta-content {
		width: 100%;
		text-align: center;
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

	.cta-primary,
	.cta-secondary {
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

		.cta-primary,
		.cta-secondary {
			width: 100%;
			max-width: 300px;
		}

		.square-wrapper-1 {
			width: 250px;
			height: 250px;
			top: 5%;
			right: -10%;
		}

		.square-wrapper-2 {
			width: 300px;
			height: 300px;
			bottom: -5%;
			left: -15%;
		}
	}
</style>
