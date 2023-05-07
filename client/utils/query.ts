import { ParsedUrlQuery } from "querystring";

export const getRawQuery = (queries: ParsedUrlQuery) => {
	const queriesStr = Object.entries(queries)
		.map(([k, v]) => `${k}=${v}`)
		.join("&");
	return `?${queriesStr}`;
};
