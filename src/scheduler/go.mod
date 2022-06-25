module github.com/shota-imoto/helixf/scheduler

replace github.com/shota-imoto/helixf/lib => ../../lib

go 1.17

require (
	github.com/robfig/cron/v3 v3.0.1
	github.com/shota-imoto/helixf/lib v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	gorm.io/driver/mysql v1.3.3 // indirect
	gorm.io/gorm v1.23.4 // indirect
)
