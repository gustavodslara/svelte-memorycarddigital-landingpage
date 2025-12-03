import { init, register, getLocaleFromNavigator } from 'svelte-i18n';

register('en', () => import('./locales/en.json'));
register('pt-BR', () => import('./locales/pt.json'));
register('pt', () => import('./locales/pt.json')); // Fallback for generic 'pt'

init({
    fallbackLocale: 'en',
    initialLocale: getLocaleFromNavigator(),
});
