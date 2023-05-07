package regular_schedule

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shota-imoto/helixf/lib/utils/schedule"
)

type RegularScheduleTemplate struct {
	Id              uint    `gorm:"primaryKey" sql:"type:uint"`
	Hour            int     `sql:"type:int" json:"hour,string"`
	Day             int     `sql:"type:int" json:"day,string"`
	Weekday         Weekday `sql:"type:int" json:"weekday"`
	Week            int     `sql:"type:int" json:"week,string"`
	Month           int     `sql:"type:int" json:"month,string"`
	LineGroupId     uint    `sql:"type:uint" json:"groupId,string"`
	RegularSchedule RegularSchedule
	CreatedAt       time.Time
}

type Weekday time.Weekday

// いる？
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (weekday Weekday) toTimeWeekday() time.Weekday {
	switch weekday {
	case Sunday:
		return time.Sunday
	case Monday:
		return time.Monday
	case Tuesday:
		return time.Tuesday
	case Wednesday:
		return time.Wednesday
	case Thursday:
		return time.Thursday
	case Friday:
		return time.Friday
	case Saturday:
		return time.Saturday
	default:
		return time.Sunday
	}
}

func (wd *Weekday) String() string {
	return wd.toTimeWeekday().String()
}

func (t RegularScheduleTemplate) MarshalJSON() ([]byte, error) {
	// Weekday以外は特に変換必要ないので何とかスマートに書けない？
	return json.Marshal(&struct {
		Id              uint   `json:"id"`
		Hour            int    `json:"hour,string"`
		Day             int    `json:"day,string"`
		Weekday         string `json:"weekday"`
		Week            int    `json:"week,string"`
		Month           int    `json:"month,string"`
		LineGroupId     uint   `json:"groupId,string"`
		RegularSchedule RegularSchedule
		CreatedAt       time.Time
	}{
		Id:              t.Id,
		Hour:            t.Hour,
		Day:             t.Day,
		Weekday:         t.Weekday.String(),
		Week:            t.Week,
		Month:           t.Month,
		LineGroupId:     t.LineGroupId,
		RegularSchedule: t.RegularSchedule,
		CreatedAt:       t.CreatedAt,
	})
}

// RegularScheduleTemplateストラクトのJson Unmarshal用
func (weekday *Weekday) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("data should be a string, got %s", data)
	}
	var wd Weekday
	switch s {
	case "Sunday":
		wd = Sunday
	case "Monday":
		wd = Monday
	case "Tuesday":
		wd = Tuesday
	case "Wednesday":
		wd = Wednesday
	case "Thursday":
		wd = Thursday
	case "Friday":
		wd = Friday
	case "Saturday":
		wd = Saturday
	default:
		return fmt.Errorf("invalid Weekday %s", s)
	}
	*weekday = wd
	return nil
}

func (template RegularScheduleTemplate) BuildSchedule(t time.Time) RegularSchedule {
	var next_schedule time.Time

	if template.Month == 0 || template.Month == int(t.Month()) {
		next_schedule = t
	} else {
		next_schedule = schedule.AfterBegginingMonth(t, template.RestMonthForNextMonth(t))
	}

	if template.Day == 0 {
		next_schedule = template.NextXthWeekday(next_schedule)
	} else {
		next_schedule = template.AfterDayInMonth(next_schedule)
	}

	next_schedule = template.SetTime(next_schedule)
	next_regular_schedule := RegularSchedule{Date: next_schedule, RegularScheduleTemplateId: template.Id}

	return next_regular_schedule
}

func (template RegularScheduleTemplate) RestMonthForNextMonth(time time.Time) int {
	if template.Month == 0 {
		return 1
	} else {
		month_diff := template.Month - int(time.Month())
		if month_diff > 0 {
			return month_diff
		} else {
			return month_diff + 12
		}
	}
}

func (template RegularScheduleTemplate) AfterDayInMonth(now time.Time) time.Time {
	diff := template.Day - now.Day()

	if diff >= 0 {
		return now.AddDate(0, 0, diff)
	} else {
		month := template.RestMonthForNextMonth(now)
		return schedule.AfterBegginingMonth(now, month).AddDate(0, 0, template.Day-1)
	}
}
func (template RegularScheduleTemplate) NextXthWeekday(time time.Time) time.Time {
	next_schedule := schedule.NextWeekday(time, template.Weekday.toTimeWeekday())

	if template.Week > 0 {
		week := schedule.WeekNumber(next_schedule)
		diff := template.Week - week

		if diff >= 0 {
			return next_schedule.AddDate(0, 0, diff*7)
		} else {
			first_weekday := schedule.NextWeekday(schedule.AfterBegginingMonth(next_schedule, template.RestMonthForNextMonth(next_schedule)), template.Weekday.toTimeWeekday())
			return first_weekday.AddDate(0, 0, (template.Week-1)*7)
		}
	} else {
		return next_schedule
	}
}

func (template RegularScheduleTemplate) SetTime(now time.Time) time.Time {
	next_schedule := now.Truncate(24 * time.Hour)

	if template.Hour > 0 {
		next_schedule = next_schedule.Add(time.Duration(template.Hour) * time.Hour)
	}
	return next_schedule
}
