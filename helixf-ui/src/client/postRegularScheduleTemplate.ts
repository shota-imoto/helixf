import { urlHost, postProps, postHeader, unauthorizedHandler } from './common'

type postRegularScheduleTemplateProps = postProps & {
	month?: string,
	week?: string,
	weekday: string,
	day?: string,
	hour?: string
}

const postRegularScheduleTemplate = async (props: postRegularScheduleTemplateProps) => {
	const body = {
		"month": props.month,
		"week": props.week,
		"weekday": props.weekday,
		"day": props.day,
		"hour": props.hour,
	}

	const data: RequestInit = {
		method: 'POST',
		mode: 'cors',
		headers: postHeader(props.authorization),
		body: JSON.stringify(body)
	}

	const response = await fetch(urlHost + 'regular_schedule_template', data)
	.then((response) => {

		if (response.status === 401) {
			unauthorizedHandler()
		}
		return response
	}).catch((error) => {
		return error
	})
	return response.json()
}

export default postRegularScheduleTemplate