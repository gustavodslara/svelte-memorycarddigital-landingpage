<script>
	import { t } from "svelte-i18n";
	import { base } from "$app/paths";
	const currentYear = new Date().getFullYear();
</script>

<footer class="footer">
	<div class="container">
		<div class="footer-content">
			<div class="footer-left">
				<div class="brand">
					<div class="logo-icon">💾</div>
					<span class="logo-text">Memory Card Digital</span>
				</div>
				<p class="tagline">
					Open source digital preservation for your most precious
					moments. Built with modern web technologies.
				</p>
				<div class="social-links">
					<a
						href="https://github.com/gustavodslara/svelte-memorycarddigital-landingpage"
						target="_blank"
						rel="noopener noreferrer"
						class="social-link"
						aria-label="GitHub"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="20"
							height="20"
							viewBox="0 0 24 24"
							fill="none"
							stroke="currentColor"
							stroke-width="2"
							stroke-linecap="round"
							stroke-linejoin="round"
							><path
								d="M9 19c-5 1.5-5-2.5-7-3m14 6v-3.87a3.37 3.37 0 0 0-.94-2.61c3.14-.35 6.44-1.54 6.44-7A5.44 5.44 0 0 0 20 4.77 5.07 5.07 0 0 0 19.91 1S18.73.65 16 2.48a13.38 13.38 0 0 0-7 0C6.27.65 5.09 1 5.09 1A5.07 5.07 0 0 0 5 4.77a5.44 5.44 0 0 0-1.5 3.78c0 5.42 3.3 6.61 6.44 7A3.37 3.37 0 0 0 9 18.13V22"
							></path></svg
						>
					</a>
					<!-- <a
						href="https://twitter.com"
						target="_blank"
						rel="noopener noreferrer"
						class="social-link"
						aria-label="Twitter"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="20"
							height="20"
							viewBox="0 0 24 24"
							fill="none"
							stroke="currentColor"
							stroke-width="2"
							stroke-linecap="round"
							stroke-linejoin="round"
							><path
								d="M22 4s-.7 2.1-2 3.4c1.6 10-9.4 17.3-18 11.6 2.2.1 4.4-.6 6-2C3 15.5.5 9.6 3 5c2.2 2.6 5.6 4.1 9 4-.9-4.2 4-6.6 7-3.8 1.1 0 3-1.2 3-1.2z"
							></path></svg
						>
					</a> -->
				</div>
			</div>

			<div class="footer-right">
				<div class="footer-column">
					<h4>{$t("footer.product")}</h4>
					<ul>
						<li>
							<a href="{base}/about" data-sveltekit-preload-data
								>{$t("footer.about")}</a
							>
						</li>
						<li>
							<a
								href="{base}/features"
								data-sveltekit-preload-data
								>{$t("footer.features")}</a
							>
						</li>
						<li>
							<a
								href="https://github.com/gustavodslara"
								target="_blank"
								rel="noopener noreferrer">Changelog</a
							>
						</li>
					</ul>
				</div>
				<div class="footer-column">
					<h4>{$t("footer.company")}</h4>
					<ul>
						<li>
							<a
								href="https://github.com/gustavodslara"
								target="_blank"
								rel="noopener noreferrer">Discussions</a
							>
						</li>
						<li>
							<a
								href="https://github.com/gustavodslara"
								target="_blank"
								rel="noopener noreferrer">Issues</a
							>
						</li>
						<li>
							<a
								href="https://github.com/gustavodslara"
								target="_blank"
								rel="noopener noreferrer">Contributing</a
							>
						</li>
					</ul>
				</div>
				<div class="footer-column">
					<h4>{$t("footer.legal")}</h4>
					<ul>
						<li>
							<a href="{base}/privacy" data-sveltekit-preload-data
								>{$t("footer.privacy")}</a
							>
						</li>
						<li>
							<a href="{base}/terms" data-sveltekit-preload-data
								>{$t("footer.terms")}</a
							>
						</li>
						<li>
							<a href="{base}/license" data-sveltekit-preload-data
								>License</a
							>
						</li>
					</ul>
				</div>
			</div>
		</div>

		<div class="footer-bottom">
			<div class="copyright">
				{$t("footer.copyright").replace("2024", String(currentYear))}
			</div>
			<div class="made-with">
				Made with <span class="heart">♥</span> and
				<a
					href="https://kit.svelte.dev"
					target="_blank"
					rel="noopener noreferrer">SvelteKit</a
				>
			</div>
		</div>
	</div>
</footer>

<style>
	.footer {
		background: #0f172a; /* Slate 900 */
		color: #94a3b8; /* Slate 400 */
		padding: 4rem 0 2rem;
		border-top: 1px solid #1e293b; /* Slate 800 */
		font-family:
			"Inter",
			system-ui,
			-apple-system,
			sans-serif;
	}

	.container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 0 1.5rem;
	}

	.footer-content {
		display: flex;
		justify-content: space-between;
		gap: 4rem;
		margin-bottom: 4rem;
	}

	.footer-left {
		max-width: 320px;
	}

	.brand {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		margin-bottom: 1rem;
		color: #f8fafc; /* Slate 50 */
	}

	.logo-icon {
		font-size: 1.5rem;
	}

	.logo-text {
		font-size: 1.25rem;
		font-weight: 700;
		letter-spacing: -0.025em;
	}

	.tagline {
		font-size: 0.95rem;
		line-height: 1.6;
		margin-bottom: 1.5rem;
	}

	.social-links {
		display: flex;
		gap: 1rem;
	}

	.social-link {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 36px;
		height: 36px;
		border-radius: 8px;
		background: #1e293b;
		color: #cbd5e1;
		transition: all 0.2s ease;
	}

	.social-link:hover {
		background: #334155;
		color: #fff;
		transform: translateY(-2px);
	}

	.footer-right {
		display: flex;
		gap: 4rem;
		flex-wrap: wrap;
	}

	.footer-column h4 {
		color: #f8fafc;
		font-size: 0.875rem;
		font-weight: 600;
		text-transform: uppercase;
		letter-spacing: 0.05em;
		margin-bottom: 1.25rem;
	}

	.footer-column ul {
		list-style: none;
		padding: 0;
		margin: 0;
	}

	.footer-column li {
		margin-bottom: 0.75rem;
	}

	.footer-column a {
		color: #94a3b8;
		text-decoration: none;
		font-size: 0.95rem;
		transition: color 0.2s ease;
	}

	.footer-column a:hover {
		color: #38bdf8; /* Sky 400 */
	}

	.footer-bottom {
		padding-top: 2rem;
		border-top: 1px solid #1e293b;
		display: flex;
		justify-content: space-between;
		align-items: center;
		font-size: 0.875rem;
	}

	.made-with {
		display: flex;
		align-items: center;
		gap: 0.25rem;
	}

	.made-with a {
		color: #cbd5e1;
		text-decoration: none;
		font-weight: 500;
	}

	.made-with a:hover {
		text-decoration: underline;
	}

	.heart {
		color: #ef4444; /* Red 500 */
	}

	@media (max-width: 768px) {
		.footer-content {
			flex-direction: column;
			gap: 3rem;
		}

		.footer-right {
			gap: 2rem;
			justify-content: space-between;
		}

		.footer-bottom {
			flex-direction: column;
			gap: 1rem;
			text-align: center;
		}
	}
</style>
