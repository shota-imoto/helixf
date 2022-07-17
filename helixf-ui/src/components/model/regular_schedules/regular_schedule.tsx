import styled from 'styled-components'

export type RegularScheduleTemplate = {
  id: number
  hour: string
  day: string
  weekday: string
  week: string
  month: string
}

const months = [
  'January',
  'February',
  'March',
  'April',
  'May',
  'June',
  'July',
  'August',
  'September',
  'October',
  'November',
  'December'
]

const getMonthName = (number: number): string => {
  return months[number - 1]
}

const numberToXth = (number: number): string => {
  switch (number % 10) {
    case 1:
      return `${number}st`
    case 2:
      return `${number}nd`
    case 3:
      return `${number}rd`
    default:
      return `${number}th`
  }
}

const TemplateLabel = (template: RegularScheduleTemplate): string => {
  let label = ''
  if (template.month !== '0') {
    label = label.concat(`every ${getMonthName(Number(template.month))} `)
  } else {
    // month 0以外なら：毎年X月
    label = label.concat('every month ')
  }
  if (template.day !== '0') {
    // day 0以外なら：X日
    label = label.concat(`${numberToXth(Number(template.day))} `)
  } else {
    if (template.week !== '0') {
      // week weekが0以外なら：X週目
      label = label.concat(`${numberToXth(Number(template.week))} `)
    } else {
      label = label.concat('every ')
    }
    // weekday dayが0なら：X曜日
    label = label.concat(`${template.weekday} `)
  }
  if (template.hour !== '0') {
    label = label.concat(`${template.hour}:00`)
  }

  return label
}

type RegularScheduleTemplatesProps = {
  regularScheduleTemplates: RegularScheduleTemplate[]
}

const RegularScheduleTemplateList = (props: RegularScheduleTemplatesProps) => {
  return (
    <>
      {props.regularScheduleTemplates.length
        ? (
            props.regularScheduleTemplates.map((template) => {
              return <TemplateLabelDiv key={template.id}>{TemplateLabel(template)}</TemplateLabelDiv>
            })
          )
        : (
        <></>
          )}
    </>
  )
}

const TemplateLabelDiv = styled.div`
  border: 1px #000000 solid;
  margin-bottom: 8px;
`

export default RegularScheduleTemplateList
