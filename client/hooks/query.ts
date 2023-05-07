import { useRouter } from "next/router";
import { useEffect } from "react";

export const GetQuery = (key: string) => {
	const router = useRouter();

	const query = router.query;
	const value = query[key];

	if (Array.isArray(value)) {
		return value.join(", ");
	} else {
		return value || "";
	}
};
