export function entries() {
	const maps = import.meta.glob('$lib/assets/maps/*/*.svg');

	const result = [];
	for (const key in maps) {
		const params = key.split('/').pop()?.split('.').at(0)?.split('-');
		if (params === undefined || params.length !== 2) {
			continue;
		}
		result.push({
			set: params[0],
			country: params[1],
			map: params[2]
		});
	}
	return result;
}

export const prerender = true;
