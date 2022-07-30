package jobs_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/attend_confirmation"
	"github.com/shota-imoto/helixf/lib/models/line_model"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule_template"
	"github.com/shota-imoto/helixf/lib/utils/line"
	"github.com/shota-imoto/helixf/scheduler/jobs"
)

func TestSendonfirmMessages(t *testing.T) {
	group := line_model.LineGroup{GroupId: "Caa245c3d70b26b44b475553ab3ed017e"}
	db.Db.Create(&group)

	template := regular_schedule_template.RegularScheduleTemplate{LineGroupId: group.Id}
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
		SentMessage bool
	}{
		{
			Description: "今日が出欠確認日より前",
			Today:       time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			ConfirmDate: time.Date(2022, 8, 2, 0, 0, 0, 0, time.UTC),
			DeadDate:    time.Date(2022, 8, 3, 0, 0, 0, 0, time.UTC),
			SentMessage: false,
		},
		{
			Description: "今日が出欠確認日より後、かつ締切日より前",
			Today:       time.Date(2022, 8, 2, 0, 0, 0, 0, time.UTC),
			ConfirmDate: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			DeadDate:    time.Date(2022, 8, 3, 0, 0, 0, 0, time.UTC),
			SentMessage: true,
		},
		{
			Description: "今日が締切日より後",
			Today:       time.Date(2022, 8, 3, 0, 0, 0, 0, time.UTC),
			ConfirmDate: time.Date(2022, 8, 1, 0, 0, 0, 0, time.UTC),
			DeadDate:    time.Date(2022, 8, 2, 0, 0, 0, 0, time.UTC),
			SentMessage: false,
		},
	}

	for _, tt := range tests {
		fmt.Println(tt.Description)
		confirm_schedule := attend_confirmation.AttendConfirmSchedule{Date: tt.ConfirmDate, DeadDate: tt.DeadDate, AttendConfirmTemplateId: confirm_template.Id, RegularScheduleId: regular_schedule.Id}
		db.Db.Create(&confirm_schedule)

		wrapper := line.DummyBotWrapper{}
		messager := line.Messager{GroupId: group.GroupId, RegularSchedule: regular_schedule, Wrapper: &wrapper}

		jobs.SendConfirmMessages(&messager, tt.Today)

		// dummy messagerにgroupIdが格納されているかをSentMessageと照らし合わせて確認
		sent_message := wrapper.PushedId != ""
		if sent_message != tt.SentMessage {
			t.Errorf("FAIL sent message, expected: %t got: %t", tt.SentMessage, sent_message)
		}
		db.Db.Where("id = ?", confirm_schedule.Id).Delete(attend_confirmation.AttendConfirmSchedule{})
	}
}
