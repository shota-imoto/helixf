import React, { useState } from 'react'
import { useCookies } from 'react-cookie'
import { useParams } from 'react-router-dom'
import Modal from 'react-modal'
import postRegularScheduleTemplate from '../../../client/regular_schedule_templates/postRegularScheduleTemplate'
import { helixfCookieName } from '../../page/authentication'
import { RegularScheduleTemplate } from './regular_schedule'

const RegularScheduleModal = (props: {
  isOpen: boolean
  setIsOpen: React.Dispatch<React.SetStateAction<boolean>>,
	templates: RegularScheduleTemplate[],
	setTemplates: React.Dispatch<React.SetStateAction<RegularScheduleTemplate[]>>
}) => {
  const weekdays: string[] = [
    'Sunday',
    'Monday',
    'Tuesday',
    'Wednesday',
    'Thursday',
    'Friday',
    'Saturday'
  ]
  const [cookies] = useCookies([helixfCookieName])
  const [month, setMonth] = useState('')
  const [week, setWeek] = useState('')
  const [weekday, setWeekday] = useState(weekdays[0])
  const [day, setDay] = useState('')
  const [hour, setHour] = useState('')
  const { id, } = useParams()

  const WeekdayForm = () => {
    return (
      <select onChange={(e) => setWeekday(e.target.value)}>
        {weekdays.map((weekday, i) => {
          return (
            <option key={'weekday_' + i} value={weekday}>
              {weekday}
            </option>
          )
        })}
      </select>
    )
  }

  const onSubmit = async (e: React.MouseEvent<HTMLInputElement, MouseEvent>) => {
    e.preventDefault()
    if (!id) return

    const response = await postRegularScheduleTemplate({
      month: Number(month).toString(),
      week: Number(week).toString(),
      weekday,
      day: Number(day).toString(),
      hour: Number(hour).toString(),
      groupId: id,
      authorization: cookies.authorization,
    }).then((response) => response.json())

		if (response.message) {
			alert('registration is failed. please, check input value.')
			console.log(`error raised: ${response.message}`)
		} else {
			props.setTemplates([...props.templates, response])
		}

		props.setIsOpen(false)

		// reset form
		setMonth('')
		setWeek('')
		setWeekday(weekdays[0])
		setDay('')
		setHour('')
  }

  return (
    <div>
      <Modal isOpen={props.isOpen}>
        <button onClick={() => props.setIsOpen(false)}>閉じる</button>
        <div>
          <b>スケジュール登録</b>
        </div>
        <div>
          <form>
            <div>
              <label>月</label>
              <input
                type="number"
                name="month"
                onChange={(e) => {
                  setMonth(e.target.value)
                }}
              ></input>
            </div>
            <div>
              <label>週</label>
              <input
                type="number"
                name="week"
                onChange={(e) => {
                  setWeek(e.target.value)
                }}
              ></input>
            </div>
            <div>
              <label>曜日</label>
              <WeekdayForm />
            </div>
            <div>
              <label>日</label>
              <input
                type="number"
                name="day"
                onChange={(e) => {
                  setDay(e.target.value)
                }}
              ></input>
            </div>
            <div>
              <label>時</label>
              <input
                type="number"
                name="hour"
                onChange={(e) => {
                  setHour(e.target.value)
                }}
              ></input>
            </div>
            <input type="submit" value="送信" onClick={(e) => onSubmit(e)} />
          </form>
        </div>
      </Modal>
    </div>
  )
}

export default RegularScheduleModal
