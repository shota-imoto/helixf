package regular_schedule_service

import (
	"fmt"

	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/helixf_user"
	"github.com/shota-imoto/helixf/lib/models/line_model"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule_template"
)

func DeleteById(id uint, user helixf_user.User) error {
	group := line_model.LineGroup{}
	err := db.Db.
		Joins("join regular_schedule_templates on regular_schedule_templates.line_group_id = line_groups.id",
			db.Db.Where(&regular_schedule_template.RegularScheduleTemplate{Id: id})).
		Joins("join line_group_user_maps on line_group_user_maps.line_group_id = line_groups.id",
			db.Db.Where(&line_model.LineGroupUserMap{UserId: user.Id})).First(&group).Error

	if err != nil {
		return err
	}
	template := regular_schedule_template.RegularScheduleTemplate{}

	err = db.Db.Delete(&template, id).Error

	if err != nil {
		return err
	}

	return nil
}

func CreateWithValidate(template regular_schedule_template.RegularScheduleTemplate) (regular_schedule_template.RegularScheduleTemplate, error) {
	template, err := Validate(template)

	if err != nil {
		return template, err
	}

	db.Db.Create(&template)
	return template, nil
}

func Validate(template regular_schedule_template.RegularScheduleTemplate) (regular_schedule_template.RegularScheduleTemplate, error) {
	if template.Hour > 23 || template.Hour < 0 {
		return template, fmt.Errorf("Invalid Hour: %d", template.Hour)
	}

	if template.Day > 31 || template.Day < 0 {
		return template, fmt.Errorf("Invalid Day: %d", template.Day)
	}

	if template.Weekday > 6 || template.Weekday < 0 {
		return template, fmt.Errorf("Invalid Weekday: %d", template.Weekday)
	}

	if template.Week > 4 || template.Week < 0 {
		return template, fmt.Errorf("Invalid Week: %d", template.Week)
	}

	if template.Month > 12 || template.Month < 0 {
		return template, fmt.Errorf("Invalid Month: %d", template.Month)
	}

	return template, nil
}
