package main

import (
	"runtime"

	"github.com/robfig/cron/v3"

	"github.com/shota-imoto/helixf/scheduler/jobs"
)

func main() {
	c := cron.New()
	c.AddFunc("* 0 * * *", jobs.GenerateNextSchedule)
	c.Start()
	runtime.Goexit()
}
