import { useCookies } from "react-cookie"
import { useLocation } from "react-router-dom"
import { backendHost } from "../utils/url"

type AuthenticationProps = {
		children?: JSX.Element;
	}

export const helixfCookieName:string = 'helixf-cookie'

const Authorization:React.FC<AuthenticationProps> = ({children}) => {
	const [_, setCookie] = useCookies([helixfCookieName]);
	const location = useLocation();

	const query = new URLSearchParams(location.search);
	const token:(string | null) = query.get('authorization')

	if (token) {
		setCookie('authorization', token)
	}

	// useCookieは呼び出し時点のcookieの状態のみgetできるようにする
	// setCookie後のcookieを取得したいので、再度useCookiesを呼び出す
	const [cookies] = useCookies([helixfCookieName]);


	if (!cookies.authorization) {
		window.location.href = backendHost + "authenticate?redirect_path=" + location.pathname.slice(1) + "&query=" + location.search;
		return null;
	}


	return (
		<>{children}</>
	)
}

export default Authorization