import type { PageServerLoad } from './$types';
import { get } from 'svelte/store';
import { PocketBaseStore } from '$lib/stores/PocketBase';

export const load: PageServerLoad = async () => {
	const search = await get(PocketBaseStore).collection('app_maps').getFullList({ sort: 'title' })
	.then((maps) => {
		const mapData = [];
		for (const map of maps) {
			const route = `/app/${map.set}/${map.country}/${map.type}${map.date ? '/' + map.date : ''}${map.variant ? '/' + map.variant : ''}`;
			const title = map.title;
			mapData.push({
				title,
				route
			});
		}
		return mapData;
	})

	return {
		post: {
			search
		}
	};
};

export const prerender = true;
