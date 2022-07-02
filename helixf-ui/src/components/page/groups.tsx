import { useState, useEffect } from "react"
import { useLocation } from 'react-router-dom'
import { useCookies } from "react-cookie"

import registerGroups from '../../client/registerGroups'
import getListGroups from '../../client/getListGroups'
import { helixfCookieName } from './authentication'
import { Group } from '../model/group'

const GroupsIndex = () => {
	const [cookies, setCookie] = useCookies([helixfCookieName]);
	const [groups, setGroups] = useState([])

	const query = new URLSearchParams(useLocation().search);

	useEffect(() => {
		const groupId = query.get('group_id')
		if (groupId !== null) {
			registerGroups({groupId: groupId, authorization: cookies.authorization})
		}
	}, [])


	useEffect(() => {
		getListGroups({authorization: cookies.authorization}).then((response) =>{
			setGroups(response.groups)
		})
	}, [])

	console.log(groups)
	return (
		<>
			{ groups.length ? <div>{GroupList(groups)}</div>:<div>groups#index not exists</div>}

		</>
	)
}

const GroupList = (groups: Group[]) => {
	return(
		<>
			{groups.map((g) => {
				return <div key={g.id}>{g.group_name}</div>
			})}
		</>
	)
}

export default GroupsIndex