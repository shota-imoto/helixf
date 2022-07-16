import { urlHost, getProps, getHeader, unauthorizedHandler } from './common'
type getListRegularScheduleProps = getProps & {
	groupId: string
}

export const getListRegularSchedule = async (props: getListRegularScheduleProps) => {
	const data: RequestInit = {
		method: 'GET',
		mode: 'cors',
		headers: getHeader(props.authorization),
	}

	return await fetch(`${urlHost}groups/${props.groupId}/regular_schedule_templates`, data).then((response) => {
		if (response.status === 401) {
			unauthorizedHandler()
		}
		return response.json()
	})
}

export default getListRegularSchedule
