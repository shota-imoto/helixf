package attend_confirmation

import (
	"time"

	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
)

type AttendConfirmTemplate struct {
	Id uint `gorm:"primaryKey"`

	// 投票日を定期スケジュールのScheduleDayAgo日前とする
	// 直前だと定期スケジュールおよび締め切り日の設定と衝突する可能性があるので3以上でなければならない
	ScheduleDayAgo            int `sql:"type:int"`
	DeadDayAfter              int
	RegularScheduleTemplateId uint
	AttendConfirmSchedule     AttendConfirmSchedule
}

// 投票日 TODO: 出欠確認なのでVoteよりもAttendanceの方が適切では？
type AttendConfirmSchedule struct {
	Id                      uint      `gorm:"primaryKey"`
	Date                    time.Time `sql:"not null;type:date"`
	DeadDate                time.Time `sql:"not null;type:date"`
	AttendConfirmTemplateId uint
	RegularScheduleId       uint
}

type AttendConfirmCondition struct {
	Id           uint `gorm:"primaryKey"`
	MaxRejection int

	// AttendConfirmTemplate.ScheduleDayAgoより小さくなければならない
	AttendConfirmTemplateId uint
}

func (template AttendConfirmTemplate) BuildSchedule(now time.Time, regular_schedule regular_schedule.RegularSchedule) AttendConfirmSchedule {
	// 定期スケジュールよりAttendConfirmTemplate.ScheduleDayAgo日前に投票日を設定
	confirm_date := regular_schedule.Date.AddDate(0, 0, -template.ScheduleDayAgo)

	confirm_schedule := AttendConfirmSchedule{AttendConfirmTemplateId: template.Id, RegularScheduleId: regular_schedule.Id}

	// 出欠確認予定日が現在より過去の場合は、今日を出欠確認日とする
	if confirm_date.Before(now) {
		confirm_schedule.Date = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	} else {
		confirm_schedule.Date = confirm_date
	}

	dead_date := confirm_date.AddDate(0, 0, template.DeadDayAfter)

	// 出欠確認日が締切日より先場合のみdead_dateを採用する
	if confirm_schedule.Date.Before(dead_date) {
		confirm_schedule.DeadDate = dead_date
	} else {
		// 締切日を過ぎているor当日の場合は、出欠確認日の翌日を締め切りとする
		confirm_schedule.DeadDate = confirm_schedule.Date.AddDate(0, 0, 1)
	}

	// 投票日を設定する
	return confirm_schedule
}
