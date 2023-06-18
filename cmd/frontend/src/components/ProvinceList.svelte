<script>
	import { onMount } from 'svelte';

	export let provinces = [];

	onMount(async () => {
		fetch('http://localhost:5000/provinces?sort_by=created_at&order_by=asc')
			.then((response) => response.json())
			.then((jsonResponse) => {
				provinces = jsonResponse.data;
			})
			.catch((error) => {
				console.log(error);
				return [];
			});
	});
</script>

<main>
	<h1 class="title">List Provinsi di <span>Indonesia</span></h1>
	<p>Total: {provinces.length}</p>
	<ul>
		{#each provinces as province (province.id)}
			<li class="link">
				<a href="/cities/{province.id}">{province.name}</a>
			</li>
		{/each}
	</ul>
</main>
