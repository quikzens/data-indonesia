<script>
	import { onMount } from 'svelte';

	export let cityId;
	export let cityName;
	export let subdistricts = [];

	onMount(async () => {
		fetch('http://localhost:5000/subdistricts?sort_by=created_at&order_by=asc&city_id=' + cityId)
			.then((response) => response.json())
			.then((jsonResponse) => {
				subdistricts = jsonResponse.data;
			})
			.catch((error) => {
				console.log(error);
				return [];
			});

		fetch('http://localhost:5000/cities/' + cityId)
			.then((response) => response.json())
			.then((jsonResponse) => {
				cityName = jsonResponse.data.name;
			})
			.catch((error) => {
				console.log(error);
				return [];
			});
	});
</script>

<main>
	<h1 class="title">List Kecamatan di <span>{cityName}</span></h1>
	<p>Total: {subdistricts.length}</p>
	<ul>
		{#each subdistricts as subdistrict (subdistrict.id)}
			<li class="link">
				<a href="/villages/{subdistrict.id}">{subdistrict.name}</a>
			</li>
		{/each}
	</ul>
</main>
