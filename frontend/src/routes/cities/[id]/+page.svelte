<script>
	import CityList from '../../../components/CityList.svelte';

	export let data;

	let provinceId = data.params.id;
</script>

<svelte:head>
	<title>List Kota</title>
</svelte:head>

<CityList {provinceId} />
