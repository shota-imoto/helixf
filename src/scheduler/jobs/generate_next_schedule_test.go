package jobs_test

import (
	"fmt"
	"testing"

	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule_template"
	"github.com/shota-imoto/helixf/scheduler/jobs"
)

func TestGenerateNextSchedule(t *testing.T) {
	var template regular_schedule_template.RegularScheduleTemplate
	var schedule regular_schedule.RegularSchedule
	var before_count, after_count int64

	fmt.Println("# スケジュールが未生成の場合")
	template = regular_schedule_template.RegularScheduleTemplate{}
	result := db.Db.Create(&template)
	if result.Error != nil {
		t.Errorf(result.Error.Error())
	}

	fmt.Println("## スケジュールが1件、生成される")
	db.Db.Model(&regular_schedule.RegularSchedule{}).Count(&before_count)
	jobs.GenerateNextSchedule()
	db.Db.Model(&regular_schedule.RegularSchedule{}).Count(&after_count)

	if after_count-before_count != 1 {
		t.Errorf("regular schedule doesn't increase by 1")
	}

	fmt.Println("## 適切なテンプレートに紐付いている")
	db.Db.Last(&schedule)
	if schedule.RegularScheduleTemplateId != template.Id {
		t.Errorf("improper template id: %v", schedule.RegularScheduleTemplateId)
	}

	fmt.Println("# スケジュールが生成済みの場合")
	fmt.Println("## スケジュールは生成されない")

	db.Db.Model(&regular_schedule.RegularSchedule{}).Count(&before_count)
	jobs.GenerateNextSchedule()
	db.Db.Model(&regular_schedule.RegularSchedule{}).Count(&after_count)

	if after_count-before_count != 0 {
		t.Errorf("regular schedule doesn't increase by 1")
	}
}
