<script>
	import { t } from "svelte-i18n";
</script>

<section class="get-started-section">
	<div class="background-gradient"></div>

	<div class="steps-container">
		<div class="step-card-wrapper">
			<div class="step-card">
				<div class="step-number">1</div>
				<h3>{$t("getStarted.step1.title")}</h3>
				<p>{$t("getStarted.step1.description")}</p>
				<a href="/download" class="card-btn"
					>{$t("getStarted.step1.button")}</a
				>
			</div>
		</div>
		<div class="step-card-wrapper">
			<div class="step-card">
				<div class="step-number">2</div>
				<h3>{$t("getStarted.step2.title")}</h3>
				<p>{$t("getStarted.step2.description")}</p>
			</div>
		</div>
		<div class="step-card-wrapper">
			<div class="step-card">
				<div class="step-number">3</div>
				<h3>{$t("getStarted.step3.title")}</h3>
				<p>{$t("getStarted.step3.description")}</p>
			</div>
		</div>
	</div>
</section>

<style>
	.get-started-section {
		height: 100vh;
		display: flex;
		align-items: center;
		justify-content: center;
		position: relative;
		overflow: hidden;
		background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
	}

	.background-gradient {
		position: absolute;
		inset: 0;
		background: radial-gradient(
			circle at 50% 50%,
			rgba(255, 255, 255, 0.1) 0%,
			transparent 50%
		);
		animation: pulse 4s ease-in-out infinite;
	}

	@keyframes pulse {
		0%,
		100% {
			opacity: 0.5;
		}
		50% {
			opacity: 1;
		}
	}

	.steps-container {
		position: relative;
		display: flex;
		gap: 2rem;
		z-index: 1;
		/* Ensure container allows absolute positioning of children if needed, 
           but here we use flex with transforms. 
           Actually, for the spread effect to work best with absolute positioning logic in GSAP, 
           we might want to rely on the wrappers being relatively positioned or just transformed.
           The current GSAP logic uses transforms, so flex layout is fine as a starting point.
        */
		align-items: center;
		justify-content: center;
	}

	/* Wrappers for the scroll animation */
	.step-card-wrapper {
		position: relative;
	}

	.step-card {
		background: rgba(255, 255, 255, 0.15);
		backdrop-filter: blur(10px);
		border: 2px solid rgba(255, 255, 255, 0.3);
		border-radius: 20px;
		padding: 3rem 2rem;
		width: 250px;
		text-align: center;
		box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
		transition: all 0.3s ease;
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.step-card:hover {
		transform: translateY(-10px) !important;
		box-shadow: 0 30px 60px rgba(0, 0, 0, 0.3);
	}

	.step-number {
		width: 60px;
		height: 60px;
		background: white;
		color: #667eea;
		font-size: 2rem;
		font-weight: 700;
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		margin: 0 auto 1.5rem;
	}

	.step-card h3 {
		color: white;
		font-size: 1.5rem;
		margin-bottom: 0.5rem;
	}

	.step-card p {
		color: rgba(255, 255, 255, 0.9);
		font-size: 1rem;
		margin-bottom: 1.5rem;
	}

	.card-btn {
		background: white;
		color: #667eea;
		border: none;
		padding: 0.8rem 1.5rem;
		border-radius: 25px;
		font-weight: 600;
		cursor: pointer;
		transition: transform 0.2s ease;
		text-decoration: none;
		display: inline-block;
	}

	.card-btn:hover {
		transform: scale(1.05);
	}

	@media (max-width: 768px) {
		.steps-container {
			flex-direction: column;
			gap: 1rem;
		}

		.step-card {
			width: 200px;
			padding: 2rem 1.5rem;
		}
	}
</style>
