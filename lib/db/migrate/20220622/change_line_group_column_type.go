package main

import (
	"github.com/shota-imoto/helixf/lib/db"
	"github.com/shota-imoto/helixf/lib/models/line_model"
)

func main() {
	db.Db.Migrator().AlterColumn(&line_model.LineGroup{}, "GroupName")
	db.Db.Migrator().AlterColumn(&line_model.LineGroup{}, "PictureUrl")
}
