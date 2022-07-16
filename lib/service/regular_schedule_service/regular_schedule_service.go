package regular_schedule_service

import (
	"fmt"

	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule_template"
)

// func GetListRegularScheduleTemplate(user helixf_user.User, group_id string) ([]regular_schedule_template.RegularScheduleTemplate, error) {
// var regular_schedule_templates []regular_schedule_template.RegularScheduleTemplate
// result := db.Db.
// }

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
