package scheduler

import (
	"log"

	"github.com/robfig/cron"
)

func Handler() {
	c := cron.New()

	c.AddFunc("@every 5m", ScheduleDrawingDate)

	log.Println("Start Cron ")
	c.Start()
}
