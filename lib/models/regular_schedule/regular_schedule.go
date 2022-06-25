package regular_schedule

import (
	"time"

	"gorm.io/gorm"
)

type RegularSchedule struct {
	gorm.Model
	Id                        uint      `gorm:"primaryKey" `
	Date                      time.Time `sql:"not null;type:date"`
	RegularScheduleTemplateId uint
	CreatedAt                 time.Time
}
