package cron

import (
	"log/slog"
	"os"
	"time"

	"github.com/linuxoid69/school-bot/internal/school"
	tg "github.com/linuxoid69/school-bot/internal/telegram"
)

const (
	oneDay   = 86400 * 1
	fiveDays = 86400 * 5
)

func TodayReport() {
	var messageError, message string

	grades, err := school.GetGrades(
		&school.Site{
			JWT:        os.Getenv("SCHOOL_JWT"),
			URL:        os.Getenv("SCHOOL_URL"),
			EucationID: os.Getenv("SCHOOL_EUCATION_ID"),
			UserAgent:  os.Getenv("SCHOOL_USER_AGENT"),
		},
	)
	if err != nil {
		slog.Warn("Error getting grades", "error", err)

		messageError = "Ошибка получения оценок"
	}

	message, err = tg.CreateMessage(grades)
	if err != nil {
		slog.Warn("Error creating message", "error", err)

		messageError = "Ошибка создания сообщения"
	}

	if messageError != "" {
		message = messageError
	}

	mesg := tg.Message{
		Text:   message,
		ChatID: os.Getenv("SCHOOL_CHAT_ID"),
		Token:  os.Getenv("SCHOOL_TOKEN"),
	}

	if err = mesg.SendGrades(); err != nil {
		slog.Warn("Error sending message", "error", err)
	}

	slog.Info("Cron task TodayReport completed")
}

func WeekReport() {
	var messageError, message string
	perionFiveDays := time.Unix(time.Now().Unix()-fiveDays, 0).Format("02.01.2006")
	perionOneDay := time.Unix(time.Now().Unix()-oneDay, 0).Format("02.01.2006")

	grades, err := school.GetGrades(
		&school.Site{
			JWT:        os.Getenv("SCHOOL_JWT"),
			URL:        os.Getenv("SCHOOL_URL"),
			EucationID: os.Getenv("SCHOOL_EUCATION_ID"),
			UserAgent:  os.Getenv("SCHOOL_USER_AGENT"),
			DateFrom:   perionFiveDays,
			DateTo:     perionOneDay,
		},
	)
	if err != nil {
		slog.Warn("Error getting grades", "error", err)

		messageError = "Ошибка получения оценок"
	}

	message, err = tg.CreateWeekReport(perionFiveDays, perionOneDay, grades)
	if err != nil {
		slog.Warn("Error creating message", "error", err)

		messageError = "Ошибка создания сообщения"
	}

	if messageError != "" {
		message = messageError
	}

	mesg := tg.Message{
		Text:   message,
		ChatID: os.Getenv("SCHOOL_CHAT_ID"),
		Token:  os.Getenv("SCHOOL_TOKEN"),
	}

	if err = mesg.SendGrades(); err != nil {
		slog.Warn("Error sending message", "error", err)
	}

	slog.Info("Cron task WeekReport completed")
}
