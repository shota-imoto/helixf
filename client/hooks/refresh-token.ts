import { useRouter } from "next/router";
import { postHeader } from "../api/post";
import { useAuthenticationState } from "../store/hooks/authentication";
import { backendHost } from "../utils/host";
import { setLocalStorage } from "../utils/local-storage";
import { getRawQuery } from "../utils/query";
import { GetQuery } from "./query";

type RefreshTokenResponse = {
	access_token: string;
	refresh_token: string;
};

export const useRefreshToken = async () => {
	const router = useRouter();
	const queryRefreshToken = GetQuery("refresh_token");
	const [{ refleshToken: stateRefreshToken }, setAuthentication] =
		useAuthenticationState();

	const refreshToken = async () => {
		console.log("refresh token called");
		if (stateRefreshToken || queryRefreshToken) {
			const body = {
				refreshToken: stateRefreshToken || queryRefreshToken,
				redirect_path: router.pathname,
				query: getRawQuery(router.query),
			};

			const data: RequestInit = {
				method: "POST",
				mode: "cors",
				headers: postHeader(),
				body: JSON.stringify(body),
			};

			const response: Response = await fetch(
				`${backendHost}/refresh_token`,
				data
			);

			if (response.status === 200) {
				const responseBody = (await response.json()) as RefreshTokenResponse;
				setAuthentication({
					authorization: responseBody.access_token,
					refleshToken: responseBody.refresh_token,
					user: { name: "" },
				});
				setLocalStorage({
					key: "authentication",
					value: responseBody.access_token,
				});
				setLocalStorage({
					key: "refreshToken",
					value: responseBody.refresh_token,
				});
			}
			router.reload();
		}
	};

	return refreshToken;
};
