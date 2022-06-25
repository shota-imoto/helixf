package reschedule_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule_template"
	"github.com/shota-imoto/helixf/lib/models/reschedule"
)

func TestBuildSchedule(t *testing.T) {
	// テストデータの準備
	var template reschedule.RescheduleTemplate
	var schedule reschedule.VoteSchedule
	today := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	regular_template := regular_schedule_template.RegularScheduleTemplate{Day: 1, Month: 2}
	db.Db.Create(&regular_template)

	built_regular_schedule := regular_template.BuildSchedule(today)
	db.Db.Create(&built_regular_schedule)

	fmt.Println("# RescheduleTemplateが当日のX日前に設定されている場合")
	fmt.Println("## 1日前")
	template = reschedule.RescheduleTemplate{Day: 1, RegularScheduleTemplateId: regular_template.Id}
	db.Db.Create(&template)
	schedule = template.BuildSchedule(today, built_regular_schedule)
	if schedule.Date != time.Date(2022, 1, 31, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", schedule.Date)
	}

	fmt.Println("## 40日前(過去の日付)")
	template = reschedule.RescheduleTemplate{Day: 40, RegularScheduleTemplateId: regular_template.Id}
	db.Db.Create(&template)
	schedule = template.BuildSchedule(today, built_regular_schedule)
	if schedule.Date != time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", schedule.Date)
	}
}
