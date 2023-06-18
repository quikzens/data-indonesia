<script>
	import { onMount } from 'svelte';

	export let provinceId;
	export let provinceName;
	export let cities = [];

	onMount(async () => {
		fetch('http://localhost:5000/cities?sort_by=created_at&order_by=asc&province_id=' + provinceId)
			.then((response) => response.json())
			.then((jsonResponse) => {
				cities = jsonResponse.data;
			})
			.catch((error) => {
				console.log(error);
				return [];
			});

		fetch('http://localhost:5000/provinces/' + provinceId)
			.then((response) => response.json())
			.then((jsonResponse) => {
				provinceName = jsonResponse.data.name;
			})
			.catch((error) => {
				console.log(error);
				return [];
			});
	});
</script>

<main>
	<h1 class="title">List Kabupaten/Kota di Provinsi <span>{provinceName}</span></h1>
	<p>Total: {cities.length}</p>
	<ul>
		{#each cities as city (city.id)}
			<li class="link">
				<a href="/subdistricts/{city.id}">{city.name}</a>
			</li>
		{/each}
	</ul>
</main>
