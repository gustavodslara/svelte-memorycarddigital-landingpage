import { init, register, getLocaleFromNavigator } from 'svelte-i18n';

register('en', () => import('./locales/en.json'));
register('en-US', () => import('./locales/en.json'));
register('pt', () => import('./locales/pt.json'));
register('pt-BR', () => import('./locales/pt.json'));
register('pt-PT', () => import('./locales/pt.json'));

init({
    fallbackLocale: 'en',
    initialLocale: getLocaleFromNavigator(),
});
