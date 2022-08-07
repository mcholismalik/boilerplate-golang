package cron

import (
	"log"
	"time"

	"github.com/mcholismalik/boilerplate-golang/pkg/constant"
	"github.com/mcholismalik/boilerplate-golang/pkg/util/env"

	"github.com/go-co-op/gocron"
)

func CheckCron() {
	log.Println("Cron running well")
}

func Init() {
	if !env.NewEnv().GetBool(constant.IS_RUN_CRON) {
		return
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	s := gocron.NewScheduler(loc)

	// second_5
	jobSec, _ := s.Every(5).Second().Do(func() {
		CheckCron()
	})
	jobSec.SingletonMode()

	s.StartAsync()
}
