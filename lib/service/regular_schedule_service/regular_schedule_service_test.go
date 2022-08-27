package regular_schedule_service_test

import (
	"fmt"
	"testing"

	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/service/regular_schedule_service"
)

func TestCreateWithValidate(t *testing.T) {
	var template, result_template regular_schedule.RegularScheduleTemplate
	var err error

	// validationのテスト
	fmt.Println("# hour validate")
	fmt.Println("hour is too big")
	template = regular_schedule.RegularScheduleTemplate{Hour: 24, Day: 0, Weekday: 0, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err == nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("hour is max limit")
	template = regular_schedule.RegularScheduleTemplate{Hour: 23, Day: 0, Weekday: 0, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err != nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("hour is min limit")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 0, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err != nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("hour is too small")
	template = regular_schedule.RegularScheduleTemplate{Hour: -1, Day: 0, Weekday: 0, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err == nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("# day validate")

	fmt.Println("day is too big")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 32, Weekday: 0, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err == nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("day is max limit")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 31, Weekday: 0, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err != nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("day is min limit")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 0, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err != nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("day is too small")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: -1, Weekday: 0, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err == nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("# weekday validate")

	fmt.Println("weekday is too big")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 7, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err == nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("saturday")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: regular_schedule.Saturday, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err != nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("sunday")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: regular_schedule.Sunday, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err != nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("weekday is too small")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: -1, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err == nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("# week validate")

	fmt.Println("week is too big")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 0, Week: 5, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err == nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("week is max limit")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 0, Week: 4, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err != nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("week is min limit")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 0, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err != nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("week is too small")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 0, Week: -1, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err == nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("# month validate")

	fmt.Println("month is too big")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 0, Week: 0, Month: 13}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err == nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("month is max limit")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 0, Week: 0, Month: 12}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err != nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("month is min limit")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 0, Week: 0, Month: 0}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err != nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}

	fmt.Println("month is too small")
	template = regular_schedule.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 0, Week: 0, Month: -1}
	result_template, err = regular_schedule_service.CreateWithValidate(template)

	if err == nil {
		t.Errorf("couldn't check invalid value: %v", result_template)
	}
}
