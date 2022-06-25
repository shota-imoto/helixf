package jobs

import (
	"time"

	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule_template"
)

func GenerateNextSchedule() {
	var templates []regular_schedule_template.RegularScheduleTemplate
	var schedules []regular_schedule.RegularSchedule
	db.Db.Table("regular_schedule_templates").Joins("left join regular_schedules on regular_schedule_templates.id = regular_schedules.regular_schedule_template_id").Where("regular_schedules.id is null").Scan(&templates)

	for _, template := range templates {
		now := time.Now()
		schedules = append(schedules, template.BuildSchedule(now))
	}

	db.Db.CreateInBatches(schedules, 100)
	return
}

// scheduleをleft joinしてschedule is nullの場合ものだけfactory実行
// has_oneを定義する
