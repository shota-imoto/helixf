import { useState, useEffect, useContext } from 'react'
import { useParams } from 'react-router-dom'
import { useCookies } from 'react-cookie'
import { helixfCookieName } from './authentication'

import getGroup from '../../client/getGroup'
import { GroupsContext } from '../../context/groups'
import { Group } from '../model/group'

const GroupPage = () => {
	const [cookies, setCookie] = useCookies([helixfCookieName]);
	const { groups, setGroups } = useContext(GroupsContext)
	const [group, setGroup] = useState<Group>({id: 0, groupId: "", group_name: ""})
	const { id } = useParams()

	useEffect(() => {
		if (!id) return
		const idGroup = groups.find((g) => g.id.toString() === id)

		if (!idGroup) {
			getGroup({groupId: id, authorization: cookies.authorization}).then((response) => {
				setGroup(response)
			})

		} else {
			setGroup(idGroup)
		}
	}, [])

	return (
		<div>
			{ id && group ?
				<div>
					<div>
						Group detail
					</div>
					<div>
						{group.group_name}
				</div>
			</div>
				:
			<div>
				Now Loading..
			</div>}
		</div>
	)
}

export default GroupPage