<script lang="ts">
	import { page } from '$app/stores';
	import '$lib/styles/global.css';
	import { loadFromJson } from '$lib/utils/loadMap';
	import { LoadedMapStore, loadMapFromURL } from '$lib/stores/LoadedMap';
	import HorizontalBattleChart from '$lib/components/charts/battlechart/BattleChart.svelte';
	import CandidateBoxContainer from '$lib/components/candidatebox/CandidateBoxContainer.svelte';
	import { loadRegionsForView } from '$lib/utils/loadRegions';
	import { applyAutoStroke, applyFastPanZoom } from '$lib/utils/applyPanZoom';
	import { browser } from '$app/environment';
	import { PocketBaseStore } from '$lib/stores/PocketBase';
	import { PUBLIC_POCKETBASE_URI } from '$env/static/public';

	$: filter = '';

	$: mapRecord = $PocketBaseStore.collection('app_maps').getFirstListItem(filter);

	$: mapFile = mapRecord.then((res) => { return fetch(`${PUBLIC_POCKETBASE_URI}/api/files/app_maps/${res.id}/${res.file}`) });

	$: map = mapFile.then((res) => { return res.text() });

	if (browser) {
		loadMapFromURL($page.url, false).then(() => {
			if ($LoadedMapStore === null) {
				return;
			}
			const { set, country, type, year, variant } = $LoadedMapStore.map;
			filter = `set = '${set}' && country = '${country}' && type = '${type}'
			&& date = '${year ? year : ''}' && variant = '${variant ? variant : ''}'`
		});
	}

	function setupMap(node: HTMLDivElement) {
		const svg = node.querySelector<SVGElement>('svg');
		if (svg !== null) {
			applyAutoStroke(svg);
			applyFastPanZoom(svg);
		}
		loadRegionsForView(node);
		if ($LoadedMapStore !== null) {
			loadFromJson($LoadedMapStore);
		}
	}
</script>

<svelte:head>
	<title>YAPms View</title>
</svelte:head>

{#if map !== undefined}
	{#await map then map}
		<div class="flex flex-col h-full p-3">
			<CandidateBoxContainer selectable={false} transitions={false} />
			<div class="grow" />
			<div use:setupMap id="map-div" class="h-4/5 overflow-hidden">
				{@html map}
			</div>
			<div class="grow" />
			<div>
				<HorizontalBattleChart transitions={false} />
			</div>
		</div>
	{/await}
{/if}
