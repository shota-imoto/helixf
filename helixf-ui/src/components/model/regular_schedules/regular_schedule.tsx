import { useState, useEffect } from "react"
import { useParams } from 'react-router-dom'
import { useCookies } from 'react-cookie'
import { helixfCookieName } from '../../page/authentication'

import getListRegularSchedule from "../../../client/getListRegularSchedule"
import { unauthorizedHandler } from "../../../client/common"

export type RegularScheduleTemplate = {
	id: number
	hour: string
	day: string
	weekday: string
	week: number
	month: string
	year: number
}

const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];

const getMonthName = (number: number): string => {
	return months[number - 1]
}

const numberToXth = (number: number): string => {
  switch (number % 10) {
    case 1:  return `${number}st`;
    case 2:  return `${number}nd`;
    case 3:  return `${number}rd`;
    default: return `${number}th`;
	}
}

// タイプ一覧
// hour 0以外なら：X時
// day 0以外なら：X日
// weekday dayが0なら：X曜日
// week weekが0以外なら：X週目
// month 0以外なら：毎年X月

const TemplateLabel = (template: RegularScheduleTemplate): string => {
	let label = ""
	console.log(template.weekday)
	if (template.month !== "0") {
		label = label.concat(`every ${getMonthName(Number(template.month))} `)
	} else {
		label = label.concat('every month ')
	}
	if (template.day !== "0") {
		label = label.concat(numberToXth(Number(template.day)))
	} else {
		if (template.week) label = label.concat(`${numberToXth(Number(template.week))} `)
		label = label.concat(template.weekday)
	}

	return label
}

type RegularScheduleTemplatesProps = {
	regularScheduleTemplates: RegularScheduleTemplate[]
}


const RegularScheduleTemplateList = (props: RegularScheduleTemplatesProps) => {

	return (
		<>
			{props.regularScheduleTemplates.map((template) => {
				return <>
					<div>
						{TemplateLabel(template)}
					</div>
				</>
			})}
		</>
	)
}

export default RegularScheduleTemplateList