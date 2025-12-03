<script>
    import { t } from "svelte-i18n";
    import Navbar from "$lib/components/Navbar.svelte";
    import Footer from "$lib/components/Footer.svelte";
</script>

<div class="page-wrapper">
    <Navbar />

    <main class="features-container">
        <div class="header">
            <h1>
                {$t("hero.backup.title")}
                <span class="highlight">{$t("hero.backup.highlight")}</span>
            </h1>
            <p class="subtitle">{$t("hero.backup.subtitle")}</p>
        </div>

        <div class="features-grid">
            <!-- Feature 1: Backup -->
            <div class="feature-card">
                <div class="icon-wrapper">
                    <i class="fa-solid fa-floppy-disk"></i>
                </div>
                <h3>{$t("hero.backup.title")}</h3>
                <p>{$t("hero.backup.subtitle")}</p>
            </div>

            <!-- Feature 2: Versioning -->
            <div class="feature-card">
                <div class="icon-wrapper">
                    <i class="fa-solid fa-code-branch"></i>
                </div>
                <h3>{$t("hero.versioning.title")}</h3>
                <p>{$t("hero.versioning.subtitle")}</p>
            </div>

            <!-- Feature 3: Cloud -->
            <div class="feature-card">
                <div class="icon-wrapper">
                    <i class="fa-solid fa-cloud"></i>
                </div>
                <h3>{$t("hero.cloud.title")}</h3>
                <p>{$t("hero.cloud.subtitle")}</p>
            </div>

            <!-- Feature 4: Security -->
            <div class="feature-card">
                <div class="icon-wrapper">
                    <i class="fa-solid fa-shield-halved"></i>
                </div>
                <h3>{$t("hero.security.title")}</h3>
                <p>{$t("hero.security.subtitle")}</p>
            </div>
        </div>
    </main>

    <Footer />
</div>

<style>
    .page-wrapper {
        min-height: 100vh;
        display: flex;
        flex-direction: column;
        background: var(--dark-bg);
    }

    .features-container {
        flex: 1;
        max-width: 1200px;
        margin: 0 auto;
        padding: 8rem 2rem 4rem;
        width: 100%;
    }

    .header {
        text-align: center;
        margin-bottom: 4rem;
    }

    h1 {
        font-size: 3rem;
        color: var(--text-light);
        margin-bottom: 1rem;
    }

    .highlight {
        background: var(--gradient);
        -webkit-background-clip: text;
        background-clip: text;
        -webkit-text-fill-color: transparent;
    }

    .subtitle {
        color: #94a3b8;
        font-size: 1.2rem;
        max-width: 600px;
        margin: 0 auto;
    }

    .features-grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
        gap: 2rem;
    }

    .feature-card {
        background: rgba(255, 255, 255, 0.03);
        border: 1px solid rgba(255, 255, 255, 0.05);
        border-radius: 24px;
        padding: 2rem;
        transition: all 0.3s ease;
    }

    .feature-card:hover {
        transform: translateY(-5px);
        background: rgba(255, 255, 255, 0.05);
        border-color: rgba(255, 255, 255, 0.1);
    }

    .icon-wrapper {
        width: 60px;
        height: 60px;
        background: rgba(99, 102, 241, 0.1);
        border-radius: 16px;
        display: flex;
        align-items: center;
        justify-content: center;
        margin-bottom: 1.5rem;
        color: var(--primary-color);
        font-size: 1.5rem;
    }

    h3 {
        color: var(--text-light);
        font-size: 1.5rem;
        margin-bottom: 1rem;
    }

    p {
        color: #94a3b8;
        line-height: 1.6;
    }

    @media (max-width: 768px) {
        h1 {
            font-size: 2.5rem;
        }

        .features-container {
            padding-top: 6rem;
        }
    }
</style>
