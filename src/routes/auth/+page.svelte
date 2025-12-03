<script>
    import { page } from "$app/stores";
    import { goto } from "$app/navigation";
    import { login } from "$lib/stores/auth";

    let isLogin = true;
    let email = "";
    let password = "";
    let name = "";
    let error = "";
    let loading = false;

    function toggleMode() {
        isLogin = !isLogin;
        error = "";
    }

    async function handleSubmit() {
        error = "";
        loading = true;
        const endpoint = isLogin ? "/api/login" : "/api/register";
        const url = `http://localhost:8000${endpoint}`;

        try {
            const payload = { email, password };
            if (!isLogin) payload.name = name;

            const response = await fetch(url, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(payload),
            });

            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.message || "Something went wrong");
            }

            if (isLogin) {
                const token = data.token;
                // Fetch user data to store in auth store
                const userRes = await fetch("http://localhost:8000/api/user", {
                    headers: { Authorization: `Bearer ${token}` },
                });
                const userData = await userRes.json();

                login(token, userData);

                const callback = $page.url.searchParams.get("callback");
                if (callback) {
                    // Deep linking redirect
                    window.location.href = `${callback}?token=${token}`;
                } else {
                    goto("/profile");
                }
            } else {
                // Registration successful, switch to login
                isLogin = true;
                error = "Account created! Please sign in."; // Using error for success msg temporarily or add success state
            }
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
        }
    }
</script>

<div class="auth-container">
    <div class="auth-card">
        <h2>{isLogin ? "Welcome Back" : "Create Account"}</h2>
        <p class="subtitle">
            {isLogin
                ? "Enter your details to access your account"
                : "Get started with your free account today"}
        </p>

        {#if error}
            <div class="error-message">
                {error}
            </div>
        {/if}

        <form on:submit|preventDefault={handleSubmit}>
            {#if !isLogin}
                <div class="form-group">
                    <label for="name">Full Name</label>
                    <div class="input-wrapper">
                        <i class="fa-solid fa-user"></i>
                        <input
                            type="text"
                            id="name"
                            bind:value={name}
                            placeholder="John Doe"
                            required
                        />
                    </div>
                </div>
            {/if}

            <div class="form-group">
                <label for="email">Email Address</label>
                <div class="input-wrapper">
                    <i class="fa-solid fa-envelope"></i>
                    <input
                        type="email"
                        id="email"
                        bind:value={email}
                        placeholder="john@example.com"
                        required
                    />
                </div>
            </div>

            <div class="form-group">
                <label for="password">Password</label>
                <div class="input-wrapper">
                    <i class="fa-solid fa-lock"></i>
                    <input
                        type="password"
                        id="password"
                        bind:value={password}
                        placeholder="••••••••"
                        required
                    />
                </div>
            </div>

            {#if isLogin}
                <div class="forgot-password">
                    <a href="/forgot-password">Forgot Password?</a>
                </div>
            {/if}

            <button type="submit" class="submit-btn" disabled={loading}>
                {loading ? "Loading..." : isLogin ? "Sign In" : "Sign Up"}
            </button>
        </form>

        <div class="toggle-mode">
            <p>
                {isLogin
                    ? "Don't have an account?"
                    : "Already have an account?"}
                <button class="link-btn" on:click={toggleMode}>
                    {isLogin ? "Sign Up" : "Sign In"}
                </button>
            </p>
        </div>
    </div>
</div>

<style>
    .auth-container {
        min-height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;
        background: var(--dark-bg);
        padding: 2rem;
        position: relative;
        overflow: hidden;
    }

    /* Background decoration */
    .auth-container::before,
    .auth-container::after {
        content: "";
        position: absolute;
        width: 300px;
        height: 300px;
        border-radius: 50%;
        filter: blur(100px);
        opacity: 0.2;
        z-index: 0;
    }

    .auth-container::before {
        background: var(--primary-color);
        top: -50px;
        left: -50px;
    }

    .auth-container::after {
        background: var(--accent-color);
        bottom: -50px;
        right: -50px;
    }

    .auth-card {
        background: rgba(255, 255, 255, 0.05);
        backdrop-filter: blur(10px);
        border: 1px solid rgba(255, 255, 255, 0.1);
        padding: 3rem;
        border-radius: 24px;
        width: 100%;
        max-width: 450px;
        position: relative;
        z-index: 1;
        box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.5);
    }

    h2 {
        color: var(--text-light);
        font-size: 2rem;
        margin-bottom: 0.5rem;
        text-align: center;
    }

    .subtitle {
        color: #94a3b8;
        text-align: center;
        margin-bottom: 2.5rem;
    }

    .error-message {
        background: rgba(239, 68, 68, 0.1);
        color: #ef4444;
        padding: 0.75rem;
        border-radius: 8px;
        margin-bottom: 1.5rem;
        text-align: center;
        font-size: 0.9rem;
    }

    .form-group {
        margin-bottom: 1.5rem;
    }

    label {
        display: block;
        color: var(--text-light);
        margin-bottom: 0.5rem;
        font-size: 0.9rem;
        font-weight: 500;
    }

    .input-wrapper {
        position: relative;
    }

    .input-wrapper i {
        position: absolute;
        left: 1rem;
        top: 50%;
        transform: translateY(-50%);
        color: #94a3b8;
        transition: color 0.3s ease;
    }

    input {
        width: 100%;
        padding: 0.75rem 1rem 0.75rem 2.75rem;
        background: rgba(15, 23, 42, 0.6);
        border: 1px solid rgba(255, 255, 255, 0.1);
        border-radius: 12px;
        color: var(--text-light);
        font-size: 1rem;
        transition: all 0.3s ease;
    }

    input:focus {
        outline: none;
        border-color: var(--primary-color);
        background: rgba(15, 23, 42, 0.8);
        box-shadow: 0 0 0 4px rgba(99, 102, 241, 0.1);
    }

    .forgot-password {
        text-align: right;
        margin-top: -0.5rem;
        margin-bottom: 1.5rem;
    }

    .forgot-password a {
        color: var(--primary-color);
        font-size: 0.9rem;
        transition: color 0.3s ease;
    }

    .forgot-password a:hover {
        color: var(--secondary-color);
    }

    .submit-btn {
        width: 100%;
        padding: 0.875rem;
        background: var(--gradient);
        color: white;
        border-radius: 12px;
        font-size: 1rem;
        font-weight: 600;
        transition:
            transform 0.2s ease,
            box-shadow 0.2s ease;
        cursor: pointer;
    }

    .submit-btn:hover:not(:disabled) {
        transform: translateY(-2px);
        box-shadow: 0 10px 20px -10px var(--primary-color);
    }

    .submit-btn:active:not(:disabled) {
        transform: translateY(0);
    }

    .submit-btn:disabled {
        opacity: 0.7;
        cursor: not-allowed;
    }

    .toggle-mode {
        margin-top: 2rem;
        text-align: center;
        color: #94a3b8;
        font-size: 0.95rem;
    }

    .link-btn {
        background: none;
        color: var(--primary-color);
        font-weight: 600;
        padding: 0;
        margin-left: 0.5rem;
        text-decoration: underline;
        text-underline-offset: 4px;
        cursor: pointer;
    }

    .link-btn:hover {
        color: var(--secondary-color);
    }

    @media (max-width: 640px) {
        .auth-container {
            padding: 1rem;
        }

        .auth-card {
            padding: 2rem;
        }
    }
</style>
