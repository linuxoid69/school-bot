package main

import (
	"time"

	"git.my-itclub.ru/bots/school/internal/checker"
	"git.my-itclub.ru/bots/school/internal/cron"
)

func main() {
	checker.CheckEnvVars()
	cron.RunTask()
	for true {
		time.Sleep(time.Second * 1)
	}
}
