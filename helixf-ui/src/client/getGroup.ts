import { urlHost, getProps, getHeader, unauthorizedHandler } from './common'

type getGroupProps = getProps & {
	groupId: string
}

const getGroup = async (props: getGroupProps) => {
	const data: RequestInit = {
		method: 'GET',
		mode: 'cors',
		headers: getHeader(props.authorization)
	}

	return fetch(`${urlHost}groups/${props.groupId}`, data).then((response) => {
		if (response.status === 401) {
			unauthorizedHandler()
		}
		return response.json()
	})
}

export default getGroup