import { useEffect } from "react"
import { useLocation } from 'react-router-dom'
import { useCookies } from "react-cookie"

import registerGroups from '../../client/registerGroups'
import { helixfCookieName } from './authentication'

const GroupsIndex = () => {
	console.log('GroupsIndex')
	const [cookies, setCookie] = useCookies([helixfCookieName]);

	const query = new URLSearchParams(useLocation().search);

	useEffect(() => {
		const groupId = query.get('group_id')
		if (groupId !== null) {
			registerGroups({groupId: groupId, authorization: cookies.authorization})
		}
	}, [])

	return (
		<div>groups#index</div>
	)
}

export default GroupsIndex