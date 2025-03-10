package cron

import (
	"log/slog"
	"os"
	"time"

	"github.com/linuxoid69/school-bot/internal/school"
	"github.com/robfig/cron/v3"
)

const (
	OneDay     = 86400 * 1
	FiveDays   = 86400 * 5
	DataFormat = "02.01.2006"
)

func RunTask() {
	l, _ := time.LoadLocation("Europe/Moscow")

	c := cron.New()
	cron.WithLocation(l)

	login := school.Login{
		Type:     school.DefaultLoginType,
		Host:     os.Getenv("SCHOOL_HOST"),
		Login:    os.Getenv("SCHOOL_LOGIN"),
		Password: os.Getenv("SCHOOL_PASSWORD"),
	}

	token, err := login.GetJWTToken()
	if err != nil {
		slog.Error("Can't get JWT token", "error", err)
	}

	site := &school.Site{
		JWT:             token,
		JournalLocation: os.Getenv("SCHOOL_HOST") + "/" + school.JournalURL,
		EucationID:      os.Getenv("SCHOOL_EUCATION_ID"),
		UserAgent:       os.Getenv("SCHOOL_USER_AGENT"),
		DateFrom:        time.Unix(time.Now().Unix()-FiveDays, 0).Format(DataFormat),
		DateTo:          time.Unix(time.Now().Unix()-OneDay, 0).Format(DataFormat),
	}

	report := Report{Type: "today"}

	_, err = c.AddFunc(os.Getenv("SCHOOL_CRON_WORK_WEEK"), func() { report.BuildReport(site) })
	if err != nil {
		slog.Warn("Error adding cron task today_report", "error", err)
	}

	report = Report{Type: "week"}

	_, err = c.AddFunc(os.Getenv("SCHOOL_CRON_WEEK_REPORT"), func() { report.BuildReport(site) })
	if err != nil {
		slog.Warn("Error adding cron task week_report", "error", err)
	}

	c.Start()
}
