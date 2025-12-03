<script>
    import { onMount } from "svelte";
    import Navbar from "$lib/components/Navbar.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import { auth } from "$lib/stores/auth";
    import { goto } from "$app/navigation";

    import { t } from "svelte-i18n";
    import SubscriptionModal from "$lib/components/SubscriptionModal.svelte";

    let user = $state({
        name: "",
        email: "",
        avatar: "",
        subscription: {
            status: "inactive",
            plan: "none",
            expiresAt: null,
        },
    });

    let backups = $state([]);
    let fileInput;
    let uploading = $state(false);
    let showModal = $state(false);

    onMount(async () => {
        const unsubscribe = auth.subscribe((value) => {
            if (!value.isAuthenticated) {
                goto("/auth");
                return;
            }
            user = {
                ...user,
                name: value.user?.email?.split("@")[0] || "User", // Fallback name
                email: value.user?.email || "",
                avatar: `https://ui-avatars.com/api/?name=${value.user?.email || "User"}&background=6366f1&color=fff`,
            };
            fetchBackups(value.token);
            fetchSubscription(value.token);
        });

        return unsubscribe;
    });

    async function fetchSubscription(token) {
        try {
            const res = await fetch("http://localhost:8000/api/subscription", {
                headers: { Authorization: `Bearer ${token}` },
            });
            if (res.ok) {
                const data = await res.json();
                user.subscription = {
                    status: data.subscription_status,
                    plan: data.subscription_plan,
                    expiresAt: data.subscription_expires_at
                        ? new Date(
                              data.subscription_expires_at,
                          ).toLocaleDateString()
                        : null,
                };
            }
        } catch (err) {
            console.error("Failed to fetch subscription", err);
        }
    }

    async function fetchBackups(token) {
        try {
            const res = await fetch("http://localhost:8000/api/saves", {
                headers: { Authorization: `Bearer ${token}` },
            });
            if (res.ok) {
                const data = await res.json();
                backups = data.map((save) => ({
                    id: save.ID,
                    game: save.game_name,
                    platform: save.platform,
                    date: new Date(save.CreatedAt).toLocaleDateString(),
                    size: formatSize(save.size),
                    image: "https://images.igdb.com/igdb/image/upload/t_cover_big/co65u5.webp", // Placeholder
                }));
            }
        } catch (err) {
            console.error("Failed to fetch backups", err);
        }
    }

    function formatSize(bytes) {
        if (bytes === 0) return "0 B";
        const k = 1024;
        const sizes = ["B", "KB", "MB", "GB"];
        const i = Math.floor(Math.log(bytes) / Math.log(k));
        return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
    }

    async function handleDownload(id, game) {
        const token = $auth.token;
        try {
            const res = await fetch(
                `http://localhost:8000/api/saves/${id}/download`,
                {
                    headers: { Authorization: `Bearer ${token}` },
                },
            );
            if (res.ok) {
                const blob = await res.blob();
                const url = window.URL.createObjectURL(blob);
                const a = document.createElement("a");
                a.href = url;
                a.download = `${game}.zip`;
                document.body.appendChild(a);
                a.click();
                window.URL.revokeObjectURL(url);
                a.remove();
            } else {
                alert("Failed to download save");
            }
        } catch (err) {
            console.error("Download error", err);
        }
    }

    function handleOpenApp(game) {
        // Deep link to open desktop app
        window.location.href = `memorycard://game/${game}`;
    }

    async function handleFileUpload(event) {
        const file = event.target.files[0];
        if (!file) return;

        uploading = true;
        const formData = new FormData();
        formData.append("file", file);
        formData.append("game_name", file.name.replace(".zip", "")); // Simple inference
        formData.append("platform", "Unknown"); // Default

        const token = $auth.token;

        try {
            const res = await fetch("http://localhost:8000/api/saves", {
                method: "POST",
                headers: { Authorization: `Bearer ${token}` },
                body: formData,
            });

            if (res.ok) {
                await fetchBackups(token);
                alert("Upload successful!");
            } else {
                alert("Upload failed");
            }
        } catch (err) {
            console.error("Upload error", err);
            alert("Upload error");
        } finally {
            uploading = false;
            fileInput.value = "";
        }
    }
</script>

<SubscriptionModal isOpen={showModal} on:close={() => (showModal = false)} />

<div class="page-wrapper">
    <Navbar />

    <main class="profile-container">
        <div class="profile-header">
            <div class="user-info">
                <img src={user.avatar} alt={user.name} class="avatar" />
                <div class="user-details">
                    <h1>{user.name}</h1>
                    <p>{user.email}</p>
                </div>
            </div>
            <div class="stats">
                <div class="stat-item">
                    <span class="stat-value">{backups.length}</span>
                    <span class="stat-label">Backups</span>
                </div>
                <div class="stat-item">
                    <span class="stat-value">2.1 MB</span>
                    <span class="stat-label">Total Size</span>
                </div>
                <div class="stat-item">
                    <span
                        class="stat-value {user.subscription.status === 'active'
                            ? 'text-green-400'
                            : 'text-zinc-400'}"
                    >
                        {user.subscription.status === "active"
                            ? $t("subscription.active")
                            : $t("subscription.inactive")}
                    </span>
                    <div class="backups-section">
                        <div class="section-header">
                            <h2>My Game Backups</h2>
                            <div class="upload-actions">
                                <input
                                    type="file"
                                    accept=".zip"
                                    style="display: none;"
                                    bind:this={fileInput}
                                    onchange={handleFileUpload}
                                />
                                <button
                                    class="btn-primary"
                                    onclick={() => fileInput.click()}
                                    disabled={uploading}
                                >
                                    <i class="fa-solid fa-upload"></i>
                                    {uploading ? "Uploading..." : "Upload Save"}
                                </button>
                            </div>
                        </div>
                        <div class="backups-list">
                            {#each backups as backup}
                                <div class="backup-item">
                                    <div class="game-info">
                                        <div class="game-icon">
                                            <i class="fa-solid fa-gamepad"></i>
                                        </div>
                                        <div class="game-details">
                                            <h3>{backup.game}</h3>
                                            <div class="meta">
                                                <span class="platform"
                                                    >{backup.platform}</span
                                                >
                                                <span class="date"
                                                    >{backup.date}</span
                                                >
                                                <span class="size"
                                                    >{backup.size}</span
                                                >
                                            </div>
                                        </div>
                                    </div>
                                    <div class="actions">
                                        <button
                                            class="btn-secondary"
                                            onclick={() =>
                                                handleDownload(
                                                    backup.id,
                                                    backup.game,
                                                )}
                                        >
                                            <i class="fa-solid fa-download"></i>
                                            Download
                                        </button>
                                        <button
                                            class="btn-primary"
                                            onclick={() =>
                                                handleOpenApp(backup.game)}
                                        >
                                            <i class="fa-solid fa-desktop"></i>
                                            Open in App
                                        </button>
                                    </div>
                                </div>
                            {/each}
                        </div>
                    </div>
                </div>
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

    .profile-container {
        flex: 1;
        max-width: 1200px;
        margin: 0 auto;
        padding: 6rem 2rem 2rem;
        width: 100%;
    }

    .profile-header {
        background: rgba(255, 255, 255, 0.05);
        backdrop-filter: blur(10px);
        border: 1px solid rgba(255, 255, 255, 0.1);
        border-radius: 24px;
        padding: 2rem;
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 3rem;
        flex-wrap: wrap;
        gap: 2rem;
    }

    .user-info {
        display: flex;
        align-items: center;
        gap: 1.5rem;
    }

    .avatar {
        width: 80px;
        height: 80px;
        border-radius: 50%;
        border: 3px solid var(--primary-color);
    }

    .user-details h1 {
        color: var(--text-light);
        font-size: 1.8rem;
        margin-bottom: 0.25rem;
    }

    .user-details p {
        color: #94a3b8;
    }

    .stats {
        display: flex;
        gap: 3rem;
    }

    .stat-item {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .stat-value {
        color: var(--primary-color);
        font-size: 1.5rem;
        font-weight: 700;
    }

    .stat-label {
        color: #94a3b8;
        font-size: 0.9rem;
    }

    .section-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 1.5rem;
    }

    .backups-section h2 {
        color: var(--text-light);
        font-size: 1.5rem;
        margin-bottom: 0;
    }

    .backups-list {
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    .backup-item {
        background: rgba(255, 255, 255, 0.03);
        border: 1px solid rgba(255, 255, 255, 0.05);
        border-radius: 16px;
        padding: 1.5rem;
        display: flex;
        justify-content: space-between;
        align-items: center;
        transition: all 0.3s ease;
    }

    .backup-item:hover {
        background: rgba(255, 255, 255, 0.05);
        border-color: rgba(255, 255, 255, 0.1);
        transform: translateY(-2px);
    }

    .game-info {
        display: flex;
        align-items: center;
        gap: 1.5rem;
    }

    .game-icon {
        width: 50px;
        height: 50px;
        background: rgba(99, 102, 241, 0.1);
        border-radius: 12px;
        display: flex;
        align-items: center;
        justify-content: center;
        color: var(--primary-color);
        font-size: 1.5rem;
    }

    .game-details h3 {
        color: var(--text-light);
        font-size: 1.1rem;
        margin-bottom: 0.25rem;
    }

    .meta {
        display: flex;
        gap: 1rem;
        font-size: 0.85rem;
        color: #94a3b8;
    }

    .platform {
        color: var(--secondary-color);
        font-weight: 600;
    }

    .actions {
        display: flex;
        gap: 1rem;
    }

    .btn-primary,
    .btn-secondary {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        padding: 0.6rem 1.2rem;
        border-radius: 8px;
        font-size: 0.9rem;
        font-weight: 500;
        transition: all 0.2s ease;
    }

    .btn-primary {
        background: var(--gradient);
        color: white;
    }

    .btn-primary:hover {
        box-shadow: 0 4px 12px rgba(99, 102, 241, 0.3);
    }

    .btn-secondary {
        background: rgba(255, 255, 255, 0.05);
        color: var(--text-light);
        border: 1px solid rgba(255, 255, 255, 0.1);
    }

    .btn-secondary:hover {
        background: rgba(255, 255, 255, 0.1);
    }

    @media (max-width: 768px) {
        .profile-header {
            flex-direction: column;
            text-align: center;
        }

        .user-info {
            flex-direction: column;
        }

        .backup-item {
            flex-direction: column;
            gap: 1.5rem;
            align-items: flex-start;
        }

        .actions {
            width: 100%;
            justify-content: space-between;
        }

        .btn-primary,
        .btn-secondary {
            flex: 1;
            justify-content: center;
        }
    }
</style>
