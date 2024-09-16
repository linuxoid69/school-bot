package main

import (
	"log/slog"
	"time"

	"git.my-itclub.ru/bots/school/internal/checker"
	"git.my-itclub.ru/bots/school/internal/cron"
)

func main() {
	slog.Info("Start bot school")

	checker.CheckEnvVars()

	cron.RunTask()

	for {
		time.Sleep(time.Millisecond * 500)
	}
}
