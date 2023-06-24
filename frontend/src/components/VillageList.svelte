<script>
	import { onMount } from 'svelte';

	export let subdistrictId;
	export let subdistrictName;
	export let villages = [];

	onMount(async () => {
		fetch(
			'http://localhost:5000/villages?sort_by=created_at&order_by=asc&subdistrict_id=' +
				subdistrictId
		)
			.then((response) => response.json())
			.then((jsonResponse) => {
				villages = jsonResponse.data;
			})
			.catch((error) => {
				console.log(error);
				return [];
			});

		fetch('http://localhost:5000/subdistricts/' + subdistrictId)
			.then((response) => response.json())
			.then((jsonResponse) => {
				subdistrictName = jsonResponse.data.name;
			})
			.catch((error) => {
				console.log(error);
				return [];
			});
	});
</script>

<main>
	<h1 class="title">List Kelurahan/Desa di Kecamatan <span>{subdistrictName}</span></h1>
	<p>Total: {villages.length}</p>
	<ul>
		{#each villages as village (village.id)}
			<li class="link">
				<span>{village.name}</span>
			</li>
		{/each}
	</ul>
</main>
