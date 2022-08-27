package jobs

import (
	"time"

	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/attend_confirmation"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
)

func GenerateNextSchedule() {
	var templates []regular_schedule.RegularScheduleTemplate
	var regular_schedules []regular_schedule.RegularSchedule
	db.Db.Table("regular_schedule_templates").Joins("left join regular_schedules on regular_schedule_templates.id = regular_schedules.regular_schedule_template_id").Where("regular_schedules.id is null").Scan(&templates)
	now := time.Now()

	for _, template := range templates {
		regular_schedules = append(regular_schedules, template.BuildSchedule(now))
	}

	db.Db.CreateInBatches(&regular_schedules, 100)

	var template_ids []uint
	for _, t := range templates {
		template_ids = append(template_ids, t.Id)
	}

	var attend_confirm_templates []attend_confirmation.AttendConfirmTemplate
	db.Db.Where("regular_schedule_template_id in (?)", template_ids).Find(&attend_confirm_templates)

	var attend_confirm_schedules []attend_confirmation.AttendConfirmSchedule
	for _, schedule := range regular_schedules {
		var attend_confirm_template attend_confirmation.AttendConfirmTemplate

		for _, template := range attend_confirm_templates {

			// templateの数が増えると重くなるかも。attend_confirm_templatesの数を途中で削ったり工夫いる？
			if schedule.RegularScheduleTemplateId == template.RegularScheduleTemplateId {
				attend_confirm_template = template
				break
			}
		}
		// TODO: templateが見つからなかった場合のハンドリング？不要？

		attend_confirm_schedules = append(attend_confirm_schedules, attend_confirm_template.BuildSchedule(now, schedule))
	}

	db.Db.CreateInBatches(attend_confirm_schedules, 100)
}
