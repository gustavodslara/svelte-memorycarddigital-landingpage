<script>
	import Hero from "$lib/components/Hero.svelte";
	import { t } from "svelte-i18n";
</script>

<svelte:head>
	<title>{$t("meta.home.title")}</title>
	<meta name="description" content={$t("meta.home.description")} />
</svelte:head>

<Hero />
