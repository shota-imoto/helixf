import { useState, useEffect, useContext } from 'react'
import { useParams } from 'react-router-dom'
import { GroupsContext } from '../../context/groups'
import { Group } from '../model/group'

const GroupPage = () => {
	const { groups, setGroups } = useContext(GroupsContext)
	const [group, setGroup] = useState<Group>({id: 0, groupId: "", group_name: ""})
	const { id } = useParams()

	useEffect(() => {
		if (!id) return
		const idGroup = groups.find((g) => g.id.toString() === id)

		if (!idGroup) {
			// get group information
		} else {
			setGroup(idGroup)
		}
	}, [])

	console.log(groups)
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