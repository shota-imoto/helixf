import { useState, useEffect, useContext } from 'react'
import { Link, useParams } from 'react-router-dom'
import { useCookies } from 'react-cookie'
import { helixfCookieName } from './authentication'

import getGroup from '../../client/getGroup'
import { GroupsContext } from '../../context/groups'
import { Group } from '../model/group'
import RegularScheduleModal from '../model/regular_schedules/regular_schedule_modal'

const GroupPage = () => {
	const [cookies, setCookie] = useCookies([helixfCookieName]);
	const { groups, setGroups } = useContext(GroupsContext)
	const [group, setGroup] = useState<Group>({id: 0, groupId: "", group_name: ""})
	const [isOpen, setIsOpen] = useState(false)
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
						<button onClick={() => setIsOpen(true)}>新規登録</button>
						<RegularScheduleModal isOpen={isOpen} setIsOpen={setIsOpen}/>
					</div>
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