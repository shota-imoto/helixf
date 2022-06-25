export const urlHost = 'http://localhost:8080/'
export const clientHost = 'http://localhost:3000/'

export type postProps = {
	authorization: string
}

export const postHeader = (authorization: string) => {
	return new Headers({
		'Access-Control-Request-Method': 'POST',
		'Access-Control-Request-Headers': 'origin, content-type, accept, access-control-request-method, authorization',
		'Origin': clientHost,
		'Content-Type': 'application/json',
		'Accept': 'application/json',
		'Authorization': authorization
	})
}

export const unauthorizedHandler = () => {
	const cookie_updated_ary:string[] = document.cookie.split(";")

	const auth_cookie:string[] | undefined = cookie_updated_ary.filter(e => !e.includes("authorization="))
	document.cookie = "authorization=;"
}