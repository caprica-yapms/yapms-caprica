<script lang="ts">
	import MapCard from '../MapCard.svelte';
	const links = [
		{
			label: 'Results',
			route: '/app/cap/parliament/20240106/results'
		},
		{
			label: 'Blank',
			route: '/app/cap/parliament/20240106/blank'
		}
	];
</script>

<MapCard name="2nd Parliamentary Election" bg="res" attribution="Photo by Anna Shvets" {links} />
