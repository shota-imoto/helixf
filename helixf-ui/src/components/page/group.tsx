import { useState, useEffect, useContext } from 'react'
import { useParams } from 'react-router-dom'
import { useCookies } from 'react-cookie'
import { helixfCookieName } from './authentication'

import getGroup from '../../client/getGroup'
import getListRegularSchedule from '../../client/getListRegularSchedule'
import { GroupsContext } from '../../context/groups'
import { Group } from '../model/group'
import RegularScheduleModal from '../model/regular_schedules/regular_schedule_modal'
import RegularScheduleTemplateList, {
  RegularScheduleTemplate
} from '../model/regular_schedules/regular_schedule'

const GroupPage = () => {
  const [cookies] = useCookies([helixfCookieName])
  const { groups, } = useContext(GroupsContext)
  const [group, setGroup] = useState<Group>({ id: 0, groupId: '', group_name: '', })
  const [templates, setTemplates] = useState<RegularScheduleTemplate[]>([])
  const [isOpen, setIsOpen] = useState(false)
  const { id, } = useParams()

  useEffect(() => {
    if (!id) return
    const idGroup = groups.find((g) => g.id.toString() === id)

    if (!idGroup) {
      getGroup({ groupId: id, authorization: cookies.authorization, }).then((response) => {
        setGroup(response)
      })
    } else {
      setGroup(idGroup)
    }

    getListRegularSchedule({ groupId: id, authorization: cookies.authorization, }).then(
      (response) => {
        setTemplates(response.regular_schedule_templates)
      }
    )
  }, [])

  return (
    <div>
      {id && group
        ? (
        <div>
          <div>
            <button onClick={() => setIsOpen(true)}>新規登録</button>
            <RegularScheduleModal isOpen={isOpen} setIsOpen={setIsOpen} templates={templates} setTemplates={setTemplates}/>
          </div>
          <div>Group detail</div>
          <div>{group.group_name}</div>
          <RegularScheduleTemplateList regularScheduleTemplates={templates} />
        </div>
          )
        : (
        <div>Now Loading..</div>
          )}
    </div>
  )
}

export default GroupPage
