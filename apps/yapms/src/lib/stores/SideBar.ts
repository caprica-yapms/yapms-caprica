import { writable } from 'svelte/store';
import { get } from 'svelte/store';
import { z } from 'zod';
import type { RecordModel } from 'pocketbase';

export const SideBarStore = writable({
	open: false,
	title: 'YAPms',
	sources: [] as string[]
});

function loadSidebarTitle(record: RecordModel) {
	const title = record.title;
	SideBarStore.set({
		...get(SideBarStore),
		title
	});
}

function loadSidebarSources(node: SVGElement) {
	let sources = [] as string[];
	const attribute = node.getAttribute('original-source') ?? '';
	try {
		sources = z.string().array().parse(JSON.parse(attribute));
	} catch (error) {
		console.error('error parsing original-source attribute:\n\n' + error);
	}
	SideBarStore.set({
		open: get(SideBarStore).open,
		title: get(SideBarStore).title,
		sources
	});
}

export { loadSidebarTitle, loadSidebarSources };
