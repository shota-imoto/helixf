package regular_schedule

import (
	"strconv"
	"time"
)

type RegularSchedule struct {
	Id                        uint      `gorm:"primaryKey" `
	Date                      time.Time `sql:"not null;type:date"`
	RegularScheduleTemplateId uint
	CreatedAt                 time.Time
}

func (s RegularSchedule) Label() string {
	month_int := int(s.Date.Month())

	return strconv.Itoa(month_int) + "/" + strconv.Itoa(s.Date.Day()) + " " + s.Date.Weekday().String()
}
