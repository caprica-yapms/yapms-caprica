<script lang="ts">
	import MapCard from '../MapCard.svelte';
	const links = [
		{
			label: 'Caprica Regions',
			route: '/app/cap/regions'
		}
	];
</script>

<MapCard name="Generic Maps" bg="gen" attribution="" {links} />
