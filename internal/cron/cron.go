package cron

import (
	"log/slog"
	"os"
	"time"

	"github.com/linuxoid69/school-bot/internal/school"
	"github.com/robfig/cron/v3"
)

const (
	ONE_DAY     = 86400 * 1
	FIVE_DAYS   = 86400 * 5
	DATA_FORMAT = "02.01.2006"
)

func RunTask() {
	l, _ := time.LoadLocation("Europe/Moscow")

	c := cron.New()
	cron.WithLocation(l)

	login := school.Login{
		Type:     school.DEFAULT_LOGIN_TYPE,
		Host:     os.Getenv("SCHOOL_HOST"),
		Login:    os.Getenv("SCHOOL_LOGIN"),
		Password: os.Getenv("SCHOOL_PASSWORD"),
	}

	token, err := login.GetJWTToken()
	if err != nil {
		slog.Error("Can't get JWT token", "error", err)
	}

	periodFiveDays := time.Unix(time.Now().Unix()-FIVE_DAYS, 0).Format(DATA_FORMAT)
	periodOneDay := time.Unix(time.Now().Unix()-ONE_DAY, 0).Format(DATA_FORMAT)

	site := &school.Site{
		JWT:             token,
		JournalLocation: os.Getenv("SCHOOL_HOST") + "/" + school.JOURNAL_URL,
		EucationID:      os.Getenv("SCHOOL_EUCATION_ID"),
		UserAgent:       os.Getenv("SCHOOL_USER_AGENT"),
		DateFrom:        periodFiveDays,
		DateTo:          periodOneDay,
	}

	_, err = c.AddFunc(os.Getenv("SCHOOL_CRON_WORK_WEEK"), func() { TodayReport(token, site) })
	if err != nil {
		slog.Warn("Error adding cron task today_report", "error", err)
	}

	_, err = c.AddFunc(os.Getenv("SCHOOL_CRON_WEEK_REPORT"), func() { WeekReport(token, site) })
	if err != nil {
		slog.Warn("Error adding cron task week_report", "error", err)
	}

	c.Start()
}
