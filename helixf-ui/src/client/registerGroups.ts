import { urlHost, postProps, postHeader, unauthorizedHandler } from './common'

type registerGroupsProps = postProps & {
  lineGroupId: string
}

const registerGroups = (props: registerGroupsProps) => {
  const body = {
    group_id: props.lineGroupId,
  }

  const data: RequestInit = {
    method: 'POST',
    mode: 'cors',
    headers: postHeader(props.authorization),
    body: JSON.stringify(body),
  }

  fetch(urlHost + 'groups/register', data)
    .then((response) => {
      if (response.status === 401) {
        unauthorizedHandler()
      }
      return response
    })
    .catch((error) => error)
}

export default registerGroups
