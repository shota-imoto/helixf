import { Router } from "@mui/icons-material";
import { useRouter } from "next/router";
import { ReactElement, useEffect } from "react";
import { GetQuery } from "../../hooks/query";
import { useAuthenticationState } from "../../store/hooks/authentication";
import { getLocalStorage, setLocalStorage } from "../../utils/local-storage";
import { Loading } from "../loading/loading";

export const helixfCookieName: string = "helixf-cookie";

type Props = {
	children: ReactElement;
};

// 期限切れトークン eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwczovL2FjY2Vzcy5saW5lLm1lIiwic3ViIjoiVWIzOTQ2YWYzMWE5ZGUwYmJiYTQ5NTJkZTFmZGFjYzIzIiwiYXVkIjoiMTY1NzA5OTkxNCIsImV4cCI6MTY3MDY2NjMwMywiaWF0IjoxNjcwNjYyNzAzLCJhbXIiOlsicHdkIl0sIm5hbWUiOiLkupXmnKwg57-U5aSqIiwicGljdHVyZSI6Imh0dHBzOi8vcHJvZmlsZS5saW5lLXNjZG4ubmV0LzBoNENqNEhLOTBhMnhYVFgtRDg1Y1VPMnNJWlFFZ1kyMGtMM2tqQ25OSU5sb3VMM3d6YnlrblhYZE9OMWw2ZVNvNVBDb2dEeWRNWWdseiJ9._EpkBt5CALUzrBB5Yqi4dx18092T698QgOdvmT81Pg4
export const Authentication = ({ children }: Props) => {
	const router = useRouter();
	const queryAuthorization = GetQuery("authorization");
	const queryRefreshToken = GetQuery("refresh_token");
	const [{ authorization, refleshToken }, setAuthentication] =
		useAuthenticationState();

	useEffect(() => {
		if (!authorization && router.isReady) {
			const localStorageAuthorization = getLocalStorage("authentication");
			const localStorageRefreshToken = getLocalStorage("refreshToken");

			if (localStorageAuthorization) {
				// ローカルストレージにトークン保存済みの場合
				setAuthentication({
					authorization: localStorageAuthorization,
					refleshToken: localStorageRefreshToken || "",
					user: { name: "" },
				});
				return;
			}

			if (queryAuthorization) {
				// 認証完了後、URLクエリでトークン取得した場合
				setAuthentication({
					authorization: queryAuthorization,
					refleshToken: queryRefreshToken,
					user: { name: "" },
				});

				setLocalStorage({ key: "authentication", value: queryAuthorization });
				setLocalStorage({ key: "refreshToken", value: queryRefreshToken });
				return;
			}

			const pathname = router.pathname;
			const queries = router.query;
			let redirectUrl = `${process.env.NEXT_PUBLIC_BACKEND_HOST}/authenticate`;

			if (!!pathname) {
				redirectUrl += `?redirect_path=${pathname}`;
				const queriesStr = Object.entries(queries)
					.map(([k, v]) => `${k}=${v}`)
					.join("&");
				if (queriesStr.length) {
					redirectUrl += `&query=?${queriesStr}`;
				}
			}
			router.push(redirectUrl);
		}
	}, [
		authorization,
		queryAuthorization,
		queryRefreshToken,
		router,
		setAuthentication,
	]);

	if (!authorization) {
		return <Loading />;
	}

	return <>{children}</>;
};
