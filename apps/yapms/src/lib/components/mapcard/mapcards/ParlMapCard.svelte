<script lang="ts">
	import MapCard from '../MapCard.svelte';
	const links = [
		{
			label: '3rd Parliament Election',
			route: '/app/cap/parliament/20240115/blank'
		},
		{
			label: 'Combined Seats',
			route: '/app/cap/parliament/20240115/combined_blank'
		}
	];
</script>

<MapCard name="Next Parliamentary Election" bg="vote" attribution="Photo by Edmond Dantès" {links} />
