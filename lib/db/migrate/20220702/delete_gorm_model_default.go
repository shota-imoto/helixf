package main

import (
	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/helixf_user"
	"github.com/shota-imoto/helixf/lib/models/line_model"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule_template"
)

func main() {
	db.Db.Migrator().DropColumn(&line_model.LineGroup{}, "deleted_at")
	db.Db.Migrator().DropColumn(&helixf_user.User{}, "deleted_at")
	db.Db.Migrator().DropColumn(&regular_schedule.RegularSchedule{}, "deleted_at")
	db.Db.Migrator().DropColumn(&regular_schedule_template.RegularScheduleTemplate{}, "deleted_at")
}
