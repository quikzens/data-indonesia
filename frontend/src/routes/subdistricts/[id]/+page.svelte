<script>
	import SubdistrictList from '../../../components/SubdistrictList.svelte';

	export let data;

	let cityId = data.params.id;
</script>

<svelte:head>
	<title>List Kecamatan</title>
</svelte:head>

<SubdistrictList {cityId} />
