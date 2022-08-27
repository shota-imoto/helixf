package attend_confirmation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/attend_confirmation"
	"github.com/shota-imoto/helixf/lib/models/line_model"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
)

func TestBuildSchedule(t *testing.T) {
	// regular_schedule.Dateは2022/2/1
	group := line_model.LineGroup{GroupId: "dummy"}
	db.Db.Create(&group)
	regular_template := regular_schedule.RegularScheduleTemplate{Day: 1, Month: 2, LineGroupId: group.Id}
	db.Db.Create(&regular_template)

	today := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	regular_schedule := regular_template.BuildSchedule(today)
	db.Db.Create(&regular_schedule)

	tests1 := []struct {
		Description    string
		ConfirmDateAgo int
		ConfirmDate    time.Time // 出欠確認予定日or今日
	}{
		{
			Description:    "出欠確認の予定日を過ぎていない",
			ConfirmDateAgo: 30,
			ConfirmDate:    time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			Description:    "出欠確認の予定の当日",
			ConfirmDateAgo: 31,
			ConfirmDate:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			Description:    "出欠確認の予定日を過ぎている",
			ConfirmDateAgo: 32,
			ConfirmDate:    time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	fmt.Println("Confirm Date Test")

	for i, tt := range tests1 {
		fmt.Printf("test case %d: %s\n", i, tt.Description)
		template := attend_confirmation.AttendConfirmTemplate{
			ScheduleDayAgo:            tt.ConfirmDateAgo,
			DeadDayAfter:              1,
			RegularScheduleTemplateId: regular_template.Id,
		}

		schedule := template.BuildSchedule(today, regular_schedule)
		if schedule.Date != tt.ConfirmDate {
			t.Errorf("FAIL expected: %v, got: %v\n", tt.ConfirmDate, schedule.Date)
		} else {
			fmt.Println("PASS")
		}
	}

	// 2/1 - 1/1 = 31
	// 32日前を締切日とした場合に出欠確認日が締切日を過ぎているテストケースを再現できるため、確認予定日を33日前とする
	CONFIRM_DATE_AGO := 33

	tests2 := []struct {
		Description  string
		DeadDayAfter int       // 出欠確認日より前or後
		DeadDate     time.Time // 締め切り予定日or明日(今日+1日)
	}{
		{
			Description:  "出欠確認日が締切日を過ぎていない",
			DeadDayAfter: 1,
			DeadDate:     time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			Description:  "出欠確認日が締切の当日",
			DeadDayAfter: 2,
			DeadDate:     time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			Description:  "出欠確認日が締切日を1日過ぎている",
			DeadDayAfter: 3,
			DeadDate:     time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			Description:  "出欠確認日が締切日を2日過ぎている",
			DeadDayAfter: 4,
			DeadDate:     time.Date(2022, 1, 3, 0, 0, 0, 0, time.UTC),
		},
	}

	fmt.Println("Dead Date Test")

	for i, tt := range tests2 {
		fmt.Printf("test case %d: %s\n", i, tt.Description)

		template := attend_confirmation.AttendConfirmTemplate{
			ScheduleDayAgo:            CONFIRM_DATE_AGO,
			DeadDayAfter:              tt.DeadDayAfter,
			RegularScheduleTemplateId: regular_template.Id,
		}

		schedule := template.BuildSchedule(today, regular_schedule)
		if schedule.DeadDate != tt.DeadDate {
			t.Errorf("failed! expected: %v, got: %v\n", tt.DeadDate, schedule.DeadDate)
		} else {
			fmt.Println("PASS")
		}
	}
}
