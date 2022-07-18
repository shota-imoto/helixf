import { urlHost, postProps, postHeader, unauthorizedHandler } from '../common'

type deleteRegularScheduleTemplateProps = postProps & {
	templateId: number
}

const deleteRegularScheduleTemplate = async (props: deleteRegularScheduleTemplateProps) => {
  const data: RequestInit = {
    method: 'DELETE',
    mode: 'cors',
    headers: postHeader(props.authorization),
  }

  const response = await fetch(`${urlHost}regular_schedule_templates/${props.templateId}`, data)
  if (response.status === 401) {
    unauthorizedHandler()
  }
  return response
}

export default deleteRegularScheduleTemplate
