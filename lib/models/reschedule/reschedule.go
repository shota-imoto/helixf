package reschedule

import (
	"time"

	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
)

type RescheduleTemplate struct {
	Id                        uint `gorm:"primaryKey"`
	Day                       int  `sql:"type:int"`
	RegularScheduleTemplateId uint
	VoteSchedule              VoteSchedule
}

type VoteSchedule struct {
	Id                   uint      `gorm:"primaryKey"`
	Date                 time.Time `sql:"not null;type:date"`
	RescheduleTemplateId uint
}

type RescheduleCondition struct {
	Id                   uint `gorm:"primaryKey"`
	MaxRejectionNumber   int
	DeadDay              int
	RescheduleTemplateId uint
}

func (template RescheduleTemplate) BuildSchedule(now time.Time, regular_schedule regular_schedule.RegularSchedule) VoteSchedule {
	date := regular_schedule.Date.AddDate(0, 0, -template.Day)
	if date.Before(now) {
		date = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	}
	vote_schedule := VoteSchedule{Date: date, RescheduleTemplateId: template.Id}

	return vote_schedule
}
