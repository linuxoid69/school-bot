package cron

import (
	"log/slog"
	"os"
	"time"

	"github.com/robfig/cron/v3"
)

func RunTask() {
	l, _ := time.LoadLocation("Europe/Moscow")

	c := cron.New()
	cron.WithLocation(l)

	_, err := c.AddFunc(os.Getenv("SCHOOL_CRON_WORK_WEEK"), TodayReport)
	if err != nil {
		slog.Warn("Error adding cron task today_report", "error", err)
	}

	_, err = c.AddFunc(os.Getenv("SCHOOL_CRON_WEEK_REPORT"), WeekReport)
	if err != nil {
		slog.Warn("Error adding cron task week_report", "error", err)
	}

	c.Start()
}
