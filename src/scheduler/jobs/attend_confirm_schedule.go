package jobs

import (
	"time"

	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/attend_confirmation"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/utils/line"
)

func SendConfirmMessagesJob() {
	client, err := line.LinebotClient()
	if err != nil {
		// TODO: エラー通知する。処理は中断する
		return
	}
	wrapper := line.LineBotWrapper{Bot: client}
	messager := line.Messager{Wrapper: &wrapper} // GroupID？？SendConfirmMessage内でセットすべきでは？
	SendConfirmMessages(&messager.Wrapper, time.Now())
}

func SendConfirmMessages(line_wrapper *line.LineWrapper, now time.Time) []error {
	// 本日分のAttendConfirmScheduleを取得
	var confirm_schedules []attend_confirmation.AttendConfirmSchedule

	err := db.Db.Where("date < ? and dead_date > ?", now, now).Find(&confirm_schedules).Error

	if err != nil {
		return []error{err}
	}
	var schedule_ids []uint

	for _, s := range confirm_schedules {
		schedule_ids = append(schedule_ids, s.RegularScheduleId)
	}

	var schedules []regular_schedule.RegularSchedule

	err = db.Db.Where("id = ?", schedule_ids).Find(&schedules).Error

	if err != nil {
		return []error{err}
	}

	// return []error{nil}
	// 出欠確認を送信する
	var errs []error

	for _, s := range schedules {
		messager := line.Messager{Wrapper: *line_wrapper, RegularSchedule: s}
		messager.SendConfirm()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
