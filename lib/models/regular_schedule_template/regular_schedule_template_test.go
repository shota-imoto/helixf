package regular_schedule_template_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule_template"
)

func TestFactorySchedule(t *testing.T) {
	var template regular_schedule_template.RegularScheduleTemplate

	fmt.Println("# every third wednesday")
	template = regular_schedule_template.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 3, Week: 2, Month: 0}

	var test_date time.Time
	var result_schedule regular_schedule.RegularSchedule
	fmt.Println("## before week in month")
	test_date = time.Date(2022, 1, 4, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 1, 12, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## before day in month")
	test_date = time.Date(2022, 1, 11, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 1, 12, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## today")
	test_date = time.Date(2022, 1, 12, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 2, 9, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## after day in month")
	test_date = time.Date(2022, 1, 13, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 2, 9, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("# every wednesday")
	template = regular_schedule_template.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 3, Week: 0, Month: 0}

	fmt.Println("## before day in week")
	test_date = time.Date(2022, 1, 4, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 1, 5, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## today in week")
	test_date = time.Date(2022, 1, 5, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 1, 12, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## after day in week")
	test_date = time.Date(2022, 1, 6, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 1, 12, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("# every day in month")
	template = regular_schedule_template.RegularScheduleTemplate{Hour: 0, Day: 15, Weekday: 0, Week: 0, Month: 0}

	fmt.Println("## before day in month")
	test_date = time.Date(2022, 1, 14, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## today in month")
	test_date = time.Date(2022, 1, 15, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 1, 15, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## after day in month")
	test_date = time.Date(2022, 1, 16, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 2, 15, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("# ??????4/15??????????????????????????????")
	template = regular_schedule_template.RegularScheduleTemplate{Hour: 0, Day: 15, Weekday: 0, Week: 0, Month: 4}

	fmt.Println("## before month")
	test_date = time.Date(2022, 3, 15, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 4, 15, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## before day in month")
	test_date = time.Date(2022, 4, 14, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 4, 15, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## ???????????????")
	test_date = time.Date(2022, 4, 15, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 4, 15, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## after day in month")
	test_date = time.Date(2022, 4, 16, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## after month")
	test_date = time.Date(2022, 5, 15, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2023, 4, 15, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("# every third wednesday in April")
	template = regular_schedule_template.RegularScheduleTemplate{Hour: 0, Day: 0, Weekday: 3, Week: 2, Month: 4}

	fmt.Println("## before month")
	test_date = time.Date(2022, 3, 13, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 4, 13, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## before day in month")
	test_date = time.Date(2022, 4, 12, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 4, 13, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## ???????????????")
	test_date = time.Date(2022, 4, 13, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2023, 4, 12, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## ?????????????????????????????????")
	test_date = time.Date(2022, 4, 14, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2023, 4, 12, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("## ???????????????")
	test_date = time.Date(2022, 5, 13, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2023, 4, 12, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	fmt.Println("# every third wednesday 13:00 in April")
	template = regular_schedule_template.RegularScheduleTemplate{Hour: 13, Day: 0, Weekday: 3, Week: 2, Month: 4}

	fmt.Println("## before month")
	test_date = time.Date(2022, 3, 13, 1, 2, 3, 4, time.UTC)

	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 4, 13, 13, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}

	// ???????????????????????????
	fmt.Println("# other error case")
	template = regular_schedule_template.RegularScheduleTemplate{Day: 1, Month: 2}
	test_date = time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	result_schedule = template.BuildSchedule(test_date)
	if result_schedule.Date != time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC) {
		t.Errorf("wrong date %v", result_schedule.Date)
	}
}
