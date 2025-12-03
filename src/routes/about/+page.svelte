<script>
	import About from "$lib/components/About.svelte";
	import { t } from "svelte-i18n";
</script>

<svelte:head>
	<title>{$t("meta.about.title")}</title>
</svelte:head>

<About />
