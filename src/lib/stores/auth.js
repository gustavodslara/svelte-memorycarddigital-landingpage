import { writable } from 'svelte/store';
import { browser } from '$app/environment';

const storedToken = browser ? localStorage.getItem('token') : null;
const storedUser = browser ? JSON.parse(localStorage.getItem('user') || 'null') : null;

export const auth = writable({
    isAuthenticated: !!storedToken,
    token: storedToken,
    user: storedUser
});

export function login(token, user) {
    if (browser) {
        localStorage.setItem('token', token);
        localStorage.setItem('user', JSON.stringify(user));
    }
    auth.set({
        isAuthenticated: true,
        token,
        user
    });
}

export function logout() {
    if (browser) {
        localStorage.removeItem('token');
        localStorage.removeItem('user');
    }
    auth.set({
        isAuthenticated: false,
        token: null,
        user: null
    });
}
