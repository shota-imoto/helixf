import React from 'react'
import deleteRegularScheduleTemplate from '../../../client/regular_schedule_templates/deleteRegularScheduleTemplate'
import { RegularScheduleTemplate } from './regular_schedule'
import { helixfCookieName } from '../../page/authentication'
import { useCookies } from 'react-cookie'

type DeleteButtonProps = {
	templateId: number
	templates: RegularScheduleTemplate[],
	setTemplates: React.Dispatch<React.SetStateAction<RegularScheduleTemplate[]>>
}

const DeleteButton = ({ templateId, templates, setTemplates }: DeleteButtonProps) => {
	const [cookies] = useCookies([helixfCookieName])

	const onClick = ({ templateId }: {templateId: number }) => {
		deleteRegularScheduleTemplate({ templateId, authorization: cookies.authorization }).then((response) => {
			if (response.status !== 204) {
				alert('Delete is failed.')
			} else {
				setTemplates(templates.filter((template) => (template.id !== templateId)))
			}
		})
	}

	return <button onClick={() => onClick({ templateId })}>Delete</button>
}

export default DeleteButton
