<script lang="ts">
	import MapCardGrid from '$lib/components/mapcard/MapCardGrid.svelte';
	import Login from '$lib/icons/Login.svelte';
	import Swatch from '$lib/icons/Swatch.svelte';
	import ThemeModal from '$lib/components/modals/thememodal/ThemeModal.svelte';
	import { AuthModalStore, ThemeModalStore } from '$lib/stores/Modals';
	import MoreMapsModal from '$lib/components/modals/moremapsmodal/MoreMapsModal.svelte';
	import MapSearch from '$lib/components/mapsearch/MapSearch.svelte';
	import ArrowUpTray from '$lib/icons/ArrowUpTray.svelte';
	import AuthModal from '$lib/components/modals/authmodal/AuthModal.svelte';
	import { PocketBaseStore } from '$lib/stores/PocketBase';
	import type { PageData } from './$types';
	import UpdatesSidebar from '$lib/components/updatessidebar/UpdatesSidebar.svelte';
	import ParlMapCard from '$lib/components/mapcard/mapcards/ParlMapCard.svelte';
	import GenMapCard from '$lib/components/mapcard/mapcards/GenMapCard.svelte';
	import SecMapCard from '$lib/components/mapcard/mapcards/SecMapCard.svelte';

	export let data: PageData;

	function openThemeModal() {
		ThemeModalStore.set({
			...$ThemeModalStore,
			open: true
		});
	}

	function openAuthModal() {
		AuthModalStore.set({
			...$AuthModalStore,
			open: true
		});
	}
</script>

<svelte:head>
	<title>Caprica YAPms</title>
	<meta name="description" content="Create and share maps for the Caprica Government Simulation." />
</svelte:head>

<div class="flex flex-col h-full overflow-hidden">
	<div class="navbar bg-base-200">
		<div class="navbar-start">
			<span class="tooltip tooltip-right" data-tip="Please use yapms.com for importing. Import functionality on this site is not up-to-date.">
				<button disabled class="btn px-8 btn-primary mr-2 hidden md:inline"
					>Import</button
				>
				<button class="btn btn-square mr-2 inline md:hidden"
					><ArrowUpTray class="h-8 m-auto" /></button
				>
			</span>
		</div>
		<div class="navbar-center">
			<h1 class="text-2xl font-bold m-auto hidden lg:inline">
				Caprica YAPms
			</h1>
			<h1 class="text-2xl font-bold m-auto inline lg:hidden">Caprica</h1>
		</div>
		<div class="navbar-end">
			<button class="btn px-8 btn-primary mr-2 hidden md:inline" on:click={openThemeModal}
				>Theme</button
			>
			<button class="btn btn-square mr-2 inline md:hidden" on:click={openThemeModal}
				><Swatch class="h-8 m-auto" /></button
			>
			<button class="btn px-8 btn-primary mr-2 hidden md:inline" on:click={openAuthModal}
				>{$PocketBaseStore.authStore.isValid ? 'Account' : 'Login'}</button
			>
			<button class="btn btn-square mr-2 inline md:hidden" on:click={openAuthModal}
				><Login class="h-8 m-auto" /></button
			>
		</div>
	</div>

	<div class="flex flex-row h-full overflow-hidden">
		<UpdatesSidebar />
		<div class="divider md:divider-horizontal ml-0 w-0 !mr-0" />
		<div class="flex-1 md:px-5 overflow-auto overflow-x-clip pb-4">
			<MapSearch data={data.post.search} />

			<MapCardGrid>
				<ParlMapCard />
				<GenMapCard />
			</MapCardGrid>


			<MapCardGrid title="Election Results">
				<SecMapCard />
			</MapCardGrid>
		</div>
	</div>
</div>

<MoreMapsModal />

<ThemeModal />

<AuthModal />
