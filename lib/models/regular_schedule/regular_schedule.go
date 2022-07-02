package regular_schedule

import (
	"time"
)

type RegularSchedule struct {
	Id                        uint      `gorm:"primaryKey" `
	Date                      time.Time `sql:"not null;type:date"`
	RegularScheduleTemplateId uint
	CreatedAt                 time.Time
}
