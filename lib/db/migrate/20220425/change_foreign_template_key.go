package main

import (
	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
)

func main() {
	db.Db.Migrator().DropColumn(&regular_schedule.RegularSchedule{}, "RegularScheduleTemplateId")
	db.Db.Migrator().AddColumn(&regular_schedule.RegularSchedule{}, "RegularScheduleTemplateId")
}
