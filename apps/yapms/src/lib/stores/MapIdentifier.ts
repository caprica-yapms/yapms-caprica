import { get, writable } from 'svelte/store';
import type { RecordModel } from 'pocketbase';

export const MapIdentifier = writable<{
	set: string;
	country: string;
	type: string;
	date: string | undefined;
	variant: string | undefined;
} | null>(null);

function loadMapIdentifier(record: RecordModel ) {
	const set = record.set;
	const country = record.country;
	const type = record.type;
	const date = record.date != '' ? record.date : undefined;
	const variant = record.variant != '' ? record.variant : undefined;

	MapIdentifier.set({
		set,
		country,
		type,
		date,
		variant,
	});

	console.log(get(MapIdentifier))
}

export { loadMapIdentifier };
