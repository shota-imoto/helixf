package db

import (
	"io/ioutil"
	"time"

	"github.com/shota-imoto/helixf/lib/models/attend_confirmation"
	"github.com/shota-imoto/helixf/lib/models/helixf_user"
	"github.com/shota-imoto/helixf/lib/models/line_model"
	"github.com/shota-imoto/helixf/lib/models/regular_schedule"
	"github.com/shota-imoto/helixf/lib/utils/helixf_env"
	"gopkg.in/yaml.v2"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// mysqlの接続
type ConnectOpenConfig struct {
	Interval int
	Retry    int
	Count    int
}

func (config ConnectOpenConfig) intervalSleep() {
	time.Sleep(time.Second * time.Duration(config.Interval))
}

func init() {
	config := ConnectOpenConfig{Interval: 5, Retry: 20, Count: 0}
	var err error

	for {
		config.intervalSleep()

		Db, err = gorm.Open(mysql.Open(dsn()), &gorm.Config{})
		if err == nil {
			break
		} else {
			if config.Count > config.Retry {
				panic(err)
			}
			config.Count++
		}
	}

	Db.AutoMigrate(
		&helixf_user.User{},
		&line_model.LineGroup{},
		&line_model.LineGroupUserMap{},
		&regular_schedule.RegularScheduleTemplate{},
		&regular_schedule.RegularSchedule{},
		&attend_confirmation.AttendConfirmTemplate{},
		&attend_confirmation.AttendConfirmSchedule{},
	)

	if helixf_env.HelixfEnv == "test" {
		Db = Db.Begin()
	}
}

type DbConfig struct {
	DbName string `yaml:"db_name"`
}

func (config DbConfig) dsn() string {
	return "root@tcp(db:3306)/" + config.DbName + "?parseTime=true"
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
