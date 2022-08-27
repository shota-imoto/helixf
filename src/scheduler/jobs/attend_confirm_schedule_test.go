package jobs_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/attend_confirmation"
	"github.com/shota-imoto/helixf/lib/models/line_model"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/utils/line"
	"github.com/shota-imoto/helixf/scheduler/jobs"
)

func TestSendonfirmMessages(t *testing.T) {
	group := line_model.LineGroup{GroupId: "Caa245c3d70b26b44b475553ab3ed017e"}
	db.Db.Create(&group)

	template := regular_schedule.RegularScheduleTemplate{LineGroupId: group.Id}
	db.Db.Create(&template)

	regular_schedule := regular_schedule.RegularSchedule{Date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), RegularScheduleTemplateId: template.Id}
	db.Db.Create(&regular_schedule)

	confirm_template := attend_confirmation.AttendConfirmTemplate{RegularScheduleTemplateId: template.Id}
	db.Db.Create(&confirm_template)

	tests := []struct {
		Description string
		Today       time.Time
		ConfirmDate time.Time
		DeadDate    time.Time
		SentCount   int
	}{
		{
			Description: "今日が出欠確認日より前",
			Today:       time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			ConfirmDate: time.Date(2022, 8, 2, 0, 0, 0, 0, time.UTC),
			DeadDate:    time.Date(2022, 8, 3, 0, 0, 0, 0, time.UTC),
			SentCount:   0,
		},
		{
			Description: "今日が出欠確認日より後、かつ締切日より前",
			Today:       time.Date(2022, 8, 2, 0, 0, 0, 0, time.UTC),
			ConfirmDate: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			DeadDate:    time.Date(2022, 8, 3, 0, 0, 0, 0, time.UTC),
			SentCount:   1,
		},
		{
			Description: "今日が締切日より後",
			Today:       time.Date(2022, 8, 3, 0, 0, 0, 0, time.UTC),
			ConfirmDate: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			DeadDate:    time.Date(2022, 8, 2, 0, 0, 0, 0, time.UTC),
			SentCount:   0,
		},
	}

	for _, tt := range tests {
		fmt.Println(tt.Description)
		confirm_schedule := attend_confirmation.AttendConfirmSchedule{Date: tt.ConfirmDate, DeadDate: tt.DeadDate, AttendConfirmTemplateId: confirm_template.Id, RegularScheduleId: regular_schedule.Id}
		db.Db.Create(&confirm_schedule)

		// wrapper := line.DummyBotWrapper{}
		// var _ line.LineWrapper = (*wrapper).(nil)
		var wrapper line.LineWrapper = &line.DummyBotWrapper{}
		// messager := line.Messager{GroupId: group.GroupId, RegularSchedule: regular_schedule, Wrapper: &wrapper}
		jobs.SendConfirmMessages(&wrapper, tt.Today)
		bot_wrapper := wrapper.(*line.DummyBotWrapper)
		// interfaceからfieldの値を取り出したい。structに変換するにはどうすれば？switch文書くしかない？

		// dummy messagerにgroupIdが格納されているかをSentMessageと照らし合わせて確認

		if bot_wrapper.CalledCount != tt.SentCount {
			t.Errorf("FAIL expected: %d, got: %d, messages: %v", tt.SentCount, bot_wrapper.CalledCount, bot_wrapper.PushedMessages)
		}
		db.Db.Where("id = ?", confirm_schedule.Id).Delete(attend_confirmation.AttendConfirmSchedule{})
	}
}
