package main

import (
	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/helixf_user"
	"github.com/shota-imoto/helixf/lib/models/line_model"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
)

func main() {
	db.Db.Migrator().DropColumn(&line_model.LineGroup{}, "deleted_at")
	db.Db.Migrator().DropColumn(&helixf_user.User{}, "deleted_at")
	db.Db.Migrator().DropColumn(&regular_schedule.RegularSchedule{}, "deleted_at")
	db.Db.Migrator().DropColumn(&regular_schedule.RegularScheduleTemplate{}, "deleted_at")
}
