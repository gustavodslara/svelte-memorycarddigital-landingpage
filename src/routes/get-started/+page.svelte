<script>
	import GetStarted from "$lib/components/GetStarted.svelte";
	import { t } from "svelte-i18n";
</script>

<svelte:head>
	<title>{$t("meta.getStarted.title")}</title>
	<meta name="description" content={$t("meta.getStarted.description")} />
</svelte:head>

<GetStarted />
