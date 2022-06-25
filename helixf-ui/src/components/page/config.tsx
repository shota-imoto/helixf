import { useState } from 'react'
import { useCookies } from "react-cookie"

import postRegularScheduleTemplate from "../../client/postRegularScheduleTemplate"
import { postProps } from "../../client/common"
import { helixfCookieName } from './authentication'
import { convertCompilerOptionsFromJson } from 'typescript'

const RegularScheduleTemplateConfig = () => {
	const weekdays:string[] = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday']
	const [cookies, setCookie] = useCookies([helixfCookieName]);
	const [month, setMonth] = useState("")
	const [week, setWeek] = useState("")
	const [weekday, setWeekday] = useState(weekdays[0])
	const [day, setDay] = useState("")
	const [hour, setHour] = useState("")

	const WeekdayForm = () => {
		return (
		<select onChange={(e) => setWeekday(e.target.value)}>
			{
				weekdays.map((weekday, i) => {
					return <option key={"weekday_" + i} value={weekday}>{weekday}</option>
				})
			}
		</select>
		)
	}

	const onSubmit = (e : React.MouseEvent<HTMLInputElement, MouseEvent>) => {
		e.preventDefault()
		const corrected_month = month === "" ? "0" : month
		const corrected_week = week === "" ? "0" : week
		const corrected_day = day === "" ? "0" : day
		const corrected_hour = hour === "" ? "0" : hour

		postRegularScheduleTemplate({month: corrected_month, week: corrected_week, weekday: weekday, day: corrected_day, hour: corrected_hour, authorization: cookies.authorization})
	}

	return (
		<div>
			<div><b>スケジュール登録</b></div>
			<div>
				<form>
					<div>
						<label>月</label>
						<input type="number" name="month" onChange={(e) => {setMonth(e.target.value)}}></input>
					</div>
					<div>
						<label>週</label>
						<input type="number" name="week" onChange={(e) => {setWeek(e.target.value)}}></input>
					</div>
					<div>
					<label>曜日</label>
						<WeekdayForm />
					</div>
					<div>
						<label>日</label>
						<input type="number" name="day" onChange={(e) => {setDay(e.target.value)}}></input>
					</div>
					<div>
						<label>時</label>
						<input type="number" name="hour" onChange={(e) => {setHour(e.target.value)}}></input>
					</div>
					<input type="submit" value="送信" onClick={(e) => onSubmit(e)}/>
				</form>
			</div>
		</div>
	)
}

export default RegularScheduleTemplateConfig