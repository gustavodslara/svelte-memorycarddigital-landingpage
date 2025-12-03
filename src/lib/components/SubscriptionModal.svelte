<script>
    import { createEventDispatcher } from "svelte";
    import { fade, scale } from "svelte/transition";
    import { t } from "svelte-i18n";

    import { auth } from "$lib/stores/auth";
    import { goto } from "$app/navigation";

    export let isOpen = false;
    const dispatch = createEventDispatcher();

    let loading = false;
    let error = "";
    let success = "";
    let selectedPlan = "monthly";

    function close() {
        dispatch("close");
        error = "";
        success = "";
    }

    async function handleSubscribe() {
        loading = true;
        error = "";
        success = "";

        try {
            const token = localStorage.getItem("token");
            if (!token) {
                error = $t("auth.loginRequired");
                loading = false;
                // Optional: redirect to login
                // goto("/auth");
                return;
            }

            const response = await fetch(
                "http://localhost:8000/api/subscription",
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: `Bearer ${token}`,
                    },
                    body: JSON.stringify({ plan: selectedPlan }),
                },
            );

            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.message || "Subscription failed");
            }

            success = $t("subscriptionModal.success");
            setTimeout(() => {
                close();
                window.location.reload(); // Reload to update UI state
            }, 1500);
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
        }
    }
</script>

{#if isOpen}
    <div
        class="modal-backdrop"
        onclick={close}
        onkeydown={(e) => e.key === "Escape" && close()}
        role="button"
        tabindex="0"
        aria-label="Close modal"
        transition:fade={{ duration: 200 }}
    >
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
        <div
            class="modal-content"
            onclick={(e) => e.stopPropagation()}
            onkeydown={(e) => e.stopPropagation()}
            role="dialog"
            aria-modal="true"
            aria-labelledby="modal-title"
            tabindex="-1"
            transition:scale={{ duration: 200, start: 0.95 }}
        >
            <button class="close-btn" onclick={close} aria-label="Close">
                <i class="fa-solid fa-times"></i>
            </button>

            <div class="modal-header">
                <div class="icon-wrapper">
                    <i class="fa-solid fa-cloud-arrow-up"></i>
                </div>
                <h2 id="modal-title">{$t("subscriptionModal.title")}</h2>

                <div class="plan-selector">
                    <button
                        class:active={selectedPlan === "monthly"}
                        onclick={() => (selectedPlan = "monthly")}
                    >
                        {$t("subscriptionModal.monthly")}
                    </button>
                    <button
                        class:active={selectedPlan === "yearly"}
                        onclick={() => (selectedPlan = "yearly")}
                    >
                        {$t("subscriptionModal.yearly")}
                    </button>
                </div>

                <p class="price">
                    {selectedPlan === "monthly" ? "R$ 9,99" : "R$ 99,99"}
                    <span class="period"
                        >/{selectedPlan === "monthly"
                            ? $t("subscriptionModal.monthly")
                            : $t("subscriptionModal.yearly")}</span
                    >
                </p>
            </div>

            <div class="modal-body">
                {#if error}
                    <div class="error-message">{error}</div>
                {/if}
                {#if success}
                    <div class="success-message">{success}</div>
                {/if}
                <p class="description">
                    {$t("hero.cloud.subtitle")}
                </p>

                <ul class="features">
                    <li>
                        <i class="fa-solid fa-check"></i>
                        <span
                            >{$t(
                                "subscriptionModal.features.cloudBackup",
                            )}</span
                        >
                    </li>
                    <li>
                        <i class="fa-solid fa-check"></i>
                        <span>{$t("subscriptionModal.features.support")}</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-check"></i>
                        <span>{$t("subscriptionModal.features.devices")}</span>
                    </li>
                    <li>
                        <i class="fa-solid fa-check"></i>
                        <span
                            >{$t("subscriptionModal.features.versioning")}</span
                        >
                    </li>
                </ul>
            </div>

            <div class="modal-footer">
                <button class="subscribe-btn" onclick={handleSubscribe}>
                    {$t("subscriptionModal.subscribe")}
                </button>
            </div>
        </div>
    </div>
{/if}

<style>
    .modal-backdrop {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, 0.6);
        backdrop-filter: blur(5px);
        z-index: 2000;
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 1rem;
    }

    .modal-content {
        background: rgba(30, 41, 59, 0.9);
        backdrop-filter: blur(20px);
        border: 1px solid rgba(255, 255, 255, 0.1);
        border-radius: 24px;
        padding: 2.5rem;
        width: 100%;
        max-width: 400px;
        position: relative;
        box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
        color: var(--text-light);
    }

    .close-btn {
        position: absolute;
        top: 1rem;
        right: 1rem;
        background: none;
        border: none;
        color: #94a3b8;
        font-size: 1.2rem;
        cursor: pointer;
        transition: color 0.2s;
        width: 32px;
        height: 32px;
        display: flex;
        align-items: center;
        justify-content: center;
        border-radius: 50%;
    }

    .close-btn:hover {
        color: var(--text-light);
        background: rgba(255, 255, 255, 0.1);
    }

    .modal-header {
        text-align: center;
        margin-bottom: 2rem;
    }

    .icon-wrapper {
        width: 64px;
        height: 64px;
        background: rgba(99, 102, 241, 0.1);
        border-radius: 20px;
        display: flex;
        align-items: center;
        justify-content: center;
        margin: 0 auto 1.5rem;
        color: var(--primary-color);
        font-size: 1.75rem;
    }

    h2 {
        font-size: 1.75rem;
        margin-bottom: 0.5rem;
        background: var(--gradient);
        -webkit-background-clip: text;
        background-clip: text;
        -webkit-text-fill-color: transparent;
    }

    .price {
        font-size: 2.5rem;
        font-weight: 700;
        color: var(--text-light);
        display: flex;
        align-items: baseline;
        justify-content: center;
        gap: 0.25rem;
    }

    .period {
        font-size: 1rem;
        color: #94a3b8;
        font-weight: 500;
    }

    .description {
        text-align: center;
        color: #94a3b8;
        margin-bottom: 2rem;
        line-height: 1.6;
    }

    .features {
        list-style: none;
        margin-bottom: 2rem;
        display: flex;
        flex-direction: column;
        gap: 0.75rem;
    }

    .features li {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        color: var(--text-light);
        font-size: 0.95rem;
    }

    .features i {
        color: var(--secondary-color);
        font-size: 0.9rem;
    }

    .subscribe-btn {
        width: 100%;
        padding: 1rem;
        background: var(--gradient);
        color: white;
        border: none;
        border-radius: 12px;
        font-size: 1.1rem;
        font-weight: 600;
        cursor: pointer;
        transition:
            transform 0.2s,
            box-shadow 0.2s;
    }

    .subscribe-btn:hover {
        transform: translateY(-2px);
        box-shadow: 0 10px 20px -10px var(--primary-color);
    }

    .subscribe-btn:active {
        transform: translateY(0);
    }

    .plan-selector {
        display: flex;
        justify-content: center;
        gap: 1rem;
        margin-bottom: 1rem;
    }

    .plan-selector button {
        background: rgba(255, 255, 255, 0.1);
        border: 1px solid rgba(255, 255, 255, 0.2);
        color: var(--text-light);
        padding: 0.5rem 1rem;
        border-radius: 20px;
        cursor: pointer;
        transition: all 0.3s;
    }

    .plan-selector button.active {
        background: var(--primary-color);
        border-color: var(--primary-color);
        color: white;
    }

    .error-message {
        color: #ef4444;
        background: rgba(239, 68, 68, 0.1);
        padding: 0.75rem;
        border-radius: 8px;
        margin-bottom: 1rem;
        text-align: center;
        font-size: 0.9rem;
    }

    .success-message {
        color: #10b981;
        background: rgba(16, 185, 129, 0.1);
        padding: 0.75rem;
        border-radius: 8px;
        margin-bottom: 1rem;
        text-align: center;
        font-size: 0.9rem;
    }
</style>
