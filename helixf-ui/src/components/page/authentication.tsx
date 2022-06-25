import { useCookies } from "react-cookie"
import { useLocation } from "react-router-dom"
import { backendHost } from "../utils/url"

type AuthenticationProps = {
		children?: JSX.Element;
	}

export const helixfCookieName:string = 'helixf-cookie'

const Authorization:React.FC<AuthenticationProps> = ({children}) => {
	const [cookies, setCookie] = useCookies([helixfCookieName]);
	const location = useLocation();

	const query = new URLSearchParams(location.search);
	const token:(string | null) = query.get('authorization')

	if (token) {
		setCookie('authorization', token)
	}

	console.log("Authorization")

	if (!cookies.authorization) {
		console.log(backendHost + "authenticate?redirect_path=" + location.pathname.slice(1) + "&query=" +location.search)
		window.location.href = backendHost + "authenticate?redirect_path=" + location.pathname.slice(1) + "&query=" + location.search;
		return null;
	}

	return (
		<>{children}</>
	)
}

export default Authorization