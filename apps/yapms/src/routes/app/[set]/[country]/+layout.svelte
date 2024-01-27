<script lang="ts">
	import { loadFromJson } from '$lib/utils/loadMap';
	import { page } from '$app/stores';
	import { MapInsetsStore } from '$lib/stores/MapInsetsStore';
	import { applyAutoStroke, applyPanZoom } from '$lib/utils/applyPanZoom';
	import { loadRegionsForApp } from '$lib/utils/loadRegions';
	import { goto } from '$app/navigation';
	import { browser } from '$app/environment';
	import { loadMapFromURL, LoadedMapStore } from '$lib/stores/LoadedMap';
	import { loadSidebarTitle, loadSidebarSources } from '$lib/stores/SideBar';
	import { loadMapIdentifier } from '$lib/stores/MapIdentifier';
	import { loadActionGroups } from '$lib/stores/ActionGroups';
	import { PocketBaseStore } from '$lib/stores/PocketBase';
	import { PUBLIC_POCKETBASE_URI } from '$env/static/public';

	$: requestedMap = $page.url.pathname.replace('/app/', '');

	$: filter = `set = '${requestedMap.split('/').at(0)}' && country = '${requestedMap.split('/').at(1)}' && type = '${requestedMap.split('/').at(2)}'
	&& date = '${requestedMap.split('/').at(3) ? requestedMap.split('/').at(3) : ''}' 
	&& variant = '${requestedMap.split('/').at(4) ? requestedMap.split('/').at(4) : ''}'`

	$: mapRecord = $PocketBaseStore.collection('app_maps').getFirstListItem(filter);

	$: mapRecord.then((res) => { 
		loadMapIdentifier(res) ;
		loadSidebarTitle(res);
	} ) ;

	$: mapFile = mapRecord.then((res) => { return fetch(`${PUBLIC_POCKETBASE_URI}/api/files/app_maps/${res.id}/${res.file}`) });

	$: map = mapFile.then((res) => { return res.text() });

	$: map.catch(() => {
		if (browser) goto('/');
	});

	function setupMap(node: HTMLDivElement) {
		const svg = node.querySelector<SVGElement>('svg');
		if (svg !== null) {
			applyPanZoom(svg);
			applyAutoStroke(svg);
			loadSidebarSources(svg);
			loadActionGroups(svg);
		}
		loadRegionsForApp(node);
		loadMapFromURL($page.url);
	}

	$: if ($LoadedMapStore) loadFromJson($LoadedMapStore);
</script>

{#await map}
	<div class="flex justify-center w-full h-full">
		<span class="loading loading-ring loading-lg text-primary"></span>
	</div>
{:then map}
	<div
		use:setupMap
		id="map-div"
		class="overflow-hidden h-full"
		class:insets-hidden={$MapInsetsStore.hidden}
	>
		{@html map}
	</div>
{/await}
