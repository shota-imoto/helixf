import { urlHost, getProps, getHeader, unauthorizedHandler } from './common'

export const getListGroups = async (props: getProps) => {
  const data: RequestInit = {
    method: 'GET',
    mode: 'cors',
    headers: getHeader(props.authorization),
  }

  return await fetch(urlHost + 'groups', data).then((response) => {
    if (response.status === 401) {
      unauthorizedHandler()
    }
    return response.json()
  })
}

export default getListGroups
