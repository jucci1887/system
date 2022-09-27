package system

import (
	"github.com/robfig/cron/v3"
	"log"
)

type quartz struct{}

var Quartz = new(quartz)

var Cron *cron.Cron

func (q *quartz) New() *quartz {
	if Cron == nil {
		Cron = cron.New(cron.WithSeconds())
	}

	return q
}

func (q *quartz) Add(task string, callback func()) *quartz {
	_, err := Cron.AddFunc(task, callback)
	if err != nil {
		log.Println("Add cron task:", task, "error:", err)
	}

	return q
}

func (q *quartz) Start() {
	Cron.Start()
	select {}
}

func (q *quartz) Stop() {
	Cron.Stop()
}
