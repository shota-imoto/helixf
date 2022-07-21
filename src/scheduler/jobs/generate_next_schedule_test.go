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
	"github.com/shota-imoto/helixf/scheduler/jobs"
)

func TestGenerateNextSchedule(t *testing.T) {
	group := line_model.LineGroup{GroupId: "dummy"}
	db.Db.Create(&group)

	// スケジュール作成済みテンプレート
	regular_template1 := regular_schedule_template.RegularScheduleTemplate{Day: 1, Month: 2, LineGroupId: group.Id}
	db.Db.Create(&regular_template1)

	attend_confirm_template1 := attend_confirmation.AttendConfirmTemplate{ScheduleDayAgo: 2, DeadDayAfter: 1, RegularScheduleTemplateId: regular_template1.Id}
	db.Db.Create(&attend_confirm_template1)

	today := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	regular_schedule1 := regular_template1.BuildSchedule(today)
	db.Db.Create(&regular_schedule1)

	attend_confirmation1 := attend_confirm_template1.BuildSchedule(today, regular_schedule1)
	db.Db.Create(&attend_confirmation1)

	// スケジュール未作成テンプレート
	regular_template2 := regular_schedule_template.RegularScheduleTemplate{Day: 1, Month: 2, LineGroupId: group.Id}
	db.Db.Create(&regular_template2)

	attend_confirm_template2 := attend_confirmation.AttendConfirmTemplate{ScheduleDayAgo: 2, DeadDayAfter: 1, RegularScheduleTemplateId: regular_template2.Id}
	db.Db.Create(&attend_confirm_template2)

	tests := []struct {
		RegularScheduleTemplate          regular_schedule_template.RegularScheduleTemplate
		AttendConfirmTemplate            attend_confirmation.AttendConfirmTemplate
		Description                      string
		RegularScheduleCountBefore       int64
		RegularScheduleCountAfter        int64
		AttendConfirmScheduleCountBefore int64
		AttendConfirmScheduleCountAfter  int64
	}{
		{
			RegularScheduleTemplate:          regular_template1,
			AttendConfirmTemplate:            attend_confirm_template1,
			Description:                      "スケジュール作成済みテンプレート",
			RegularScheduleCountBefore:       1,
			RegularScheduleCountAfter:        1,
			AttendConfirmScheduleCountBefore: 1,
			AttendConfirmScheduleCountAfter:  1,
		},
		{
			RegularScheduleTemplate:          regular_template2,
			AttendConfirmTemplate:            attend_confirm_template2,
			Description:                      "スケジュール未作成テンプレート",
			RegularScheduleCountBefore:       0,
			RegularScheduleCountAfter:        1,
			AttendConfirmScheduleCountBefore: 0,
			AttendConfirmScheduleCountAfter:  1,
		},
	}

	fmt.Println("関数実行前の確認")
	for _, tt := range tests {
		fmt.Println(tt.Description)
		var count int64
		db.Db.Model(regular_schedule.RegularSchedule{}).Where("regular_schedule_template_id = ?", tt.RegularScheduleTemplate.Id).Count(&count)

		fmt.Println("RegularScheduleCountBefore")
		if count != tt.RegularScheduleCountBefore {
			t.Errorf("FAIL expected: %d, got: %d", tt.RegularScheduleCountBefore, count)
		}

		db.Db.Model(attend_confirmation.AttendConfirmSchedule{}).Where("attend_confirm_template_id = ?", tt.AttendConfirmTemplate.Id).Count(&count)

		fmt.Println("AttendConfirmScheduleCountBefore")
		if count != tt.AttendConfirmScheduleCountBefore {
			t.Errorf("FAIL expected: %d, got: %d", tt.AttendConfirmScheduleCountBefore, count)
		}
	}

	// 関数実行
	jobs.GenerateNextSchedule()

	fmt.Println("関数実行後の確認")
	for _, tt := range tests {
		fmt.Println(tt.Description)
		var count int64
		db.Db.Model(regular_schedule.RegularSchedule{}).Where("regular_schedule_template_id = ?", tt.RegularScheduleTemplate.Id).Count(&count)

		fmt.Println("RegularScheduleCountAfter")
		if count != tt.RegularScheduleCountAfter {
			t.Errorf("FAIL expected: %d, got: %d", tt.RegularScheduleCountAfter, count)
		}

		db.Db.Model(attend_confirmation.AttendConfirmSchedule{}).Where("attend_confirm_template_id = ?", tt.AttendConfirmTemplate.Id).Count(&count)

		fmt.Println("AttendConfirmScheduleCountAfter")
		if count != tt.AttendConfirmScheduleCountAfter {
			t.Errorf("FAIL expected: %d, got: %d", tt.AttendConfirmScheduleCountAfter, count)
		}
	}
}
