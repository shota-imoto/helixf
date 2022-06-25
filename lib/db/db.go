package db

import (
	"io/ioutil"

	"github.com/shota-imoto/helixf/lib/models/helixf_user"
	"github.com/shota-imoto/helixf/lib/models/line_model"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule_template"
	"github.com/shota-imoto/helixf/lib/models/reschedule"
	"github.com/shota-imoto/helixf/lib/utils/helixf_env"
	"gopkg.in/yaml.v2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open(mysql.Open(dsn()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if helixf_env.HelixfEnv == "test" {
		Db = Db.Begin()
	}
	Db.AutoMigrate(
		&regular_schedule_template.RegularScheduleTemplate{},
		&regular_schedule.RegularSchedule{},
		&reschedule.RescheduleTemplate{},
		&helixf_user.User{},
		&line_model.LineGroup{},
		&line_model.LineGroupUserMap{},
	)
}

type DbConfig struct {
	DbName string `yaml:"db_name"`
}

func (config DbConfig) dsn() string {
	return "helixf:helixf@tcp(localhost:3306)/" + config.DbName + "?parseTime=true"
}

func dsn() string {
	var err error
	buf, err := ioutil.ReadFile(helixf_env.RootPath + "/lib/config/db/" + helixf_env.HelixfEnv + ".yml")
	if err != nil {
		panic(err)
	}

	db_config := DbConfig{}
	err = yaml.Unmarshal(buf, &db_config)

	if err != nil {
		panic(err)
	}

	return db_config.dsn()
}
