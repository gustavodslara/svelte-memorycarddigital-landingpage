<footer class="footer">
	<div class="container">
		<div class="footer-content">
			<div class="footer-column">
				<h3 class="footer-logo">Memory Card Digital</h3>
				<p class="footer-description">
					Preserving your precious memories in beautiful digital cards.
				</p>
				<div class="social-links">
					<a href="https://facebook.com" target="_blank" rel="noopener noreferrer" class="social-link">📘</a>
					<a href="https://instagram.com" target="_blank" rel="noopener noreferrer" class="social-link">📷</a>
					<a href="https://twitter.com" target="_blank" rel="noopener noreferrer" class="social-link">🐦</a>
					<a href="https://linkedin.com" target="_blank" rel="noopener noreferrer" class="social-link">💼</a>
				</div>
			</div>

			<div class="footer-column">
				<h4 class="footer-heading">Product</h4>
				<ul class="footer-links">
					<li><a href="/features" data-sveltekit-preload-data>Features</a></li>
					<li><a href="/about" data-sveltekit-preload-data>About</a></li>
					<li><a href="/contact" data-sveltekit-preload-data>Contact</a></li>
				</ul>
			</div>

			<div class="footer-column">
				<h4 class="footer-heading">Company</h4>
				<ul class="footer-links">
					<li><a href="/about" data-sveltekit-preload-data>About Us</a></li>
					<li><a href="/contact" data-sveltekit-preload-data>Contact</a></li>
				</ul>
			</div>

			<div class="footer-column">
				<h4 class="footer-heading">Legal</h4>
				<ul class="footer-links">
					<li><a href="/contact" data-sveltekit-preload-data>Privacy Policy</a></li>
					<li><a href="/contact" data-sveltekit-preload-data>Terms of Service</a></li>
				</ul>
			</div>
		</div>

		<div class="footer-bottom">
			<p>&copy; 2025 Memory Card Digital. All rights reserved.</p>
		</div>
	</div>
</footer>

<style>
	.footer {
		background: var(--dark-bg);
		color: var(--text-light);
		padding: 4rem 0 2rem;
	}

	.container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 0 2rem;
	}

	.footer-content {
		display: grid;
		grid-template-columns: 2fr 1fr 1fr 1fr;
		gap: 3rem;
		margin-bottom: 3rem;
	}

	.footer-logo {
		font-size: 1.5rem;
		font-weight: 700;
		background: var(--gradient);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
		margin-bottom: 1rem;
	}

	.footer-description {
		color: #94a3b8;
		margin-bottom: 1.5rem;
		line-height: 1.6;
	}

	.social-links {
		display: flex;
		gap: 1rem;
	}

	.social-link {
		width: 40px;
		height: 40px;
		display: flex;
		align-items: center;
		justify-content: center;
		background: rgba(255, 255, 255, 0.1);
		border-radius: 50%;
		font-size: 1.2rem;
		transition: all 0.3s ease;
	}

	.social-link:hover {
		background: var(--primary-color);
		transform: translateY(-3px);
	}

	.footer-heading {
		font-size: 1.1rem;
		margin-bottom: 1.5rem;
		color: white;
	}

	.footer-links {
		list-style: none;
	}

	.footer-links li {
		margin-bottom: 0.75rem;
	}

	.footer-links a {
		color: #94a3b8;
		transition: color 0.3s ease;
		text-decoration: none;
	}

	.footer-links a:hover {
		color: var(--primary-color);
	}

	.footer-links button {
		color: #94a3b8;
		transition: color 0.3s ease;
		background: none;
		border: none;
		cursor: pointer;
		padding: 0;
		font-size: inherit;
		text-align: left;
	}

	.footer-links button:hover {
		color: var(--primary-color);
	}

	.footer-bottom {
		text-align: center;
		padding-top: 2rem;
		border-top: 1px solid rgba(255, 255, 255, 0.1);
		color: #64748b;
	}

	@media (max-width: 768px) {
		.footer-content {
			grid-template-columns: 1fr;
			gap: 2rem;
		}
	}
</style>
