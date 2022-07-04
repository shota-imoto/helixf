import { useState, useEffect, useContext } from "react"
import { Link, useParams, useLocation } from 'react-router-dom'
import { useCookies } from "react-cookie"
import { GroupsContext } from '../../context/groups'

import registerGroups from '../../client/registerGroups'
import getListGroups from '../../client/getListGroups'
import { helixfCookieName } from './authentication'
import { Group } from '../model/group'
import { frontendHost } from '../utils/url'

const GroupsIndex = () => {
	const [cookies, setCookie] = useCookies([helixfCookieName]);
	const { groups, setGroups } = useContext(GroupsContext)


	const query = new URLSearchParams(useLocation().search);

	useEffect(() => {
		const groupId = query.get('group_id')
		if (groupId !== null) {
			registerGroups({lineGroupId: groupId, authorization: cookies.authorization})
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
				return <div key={g.id}><Link to={`/groups/${g.id}`}>{g.group_name} </Link></div>
			})}
		</>
	)
}

export default GroupsIndex