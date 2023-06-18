<script>
	import VillageList from '../../../components/VillageList.svelte';

	export let data;

	let subdistrictId = data.params.id;
</script>

<svelte:head>
	<title>List Desa</title>
</svelte:head>

<VillageList {subdistrictId} />
