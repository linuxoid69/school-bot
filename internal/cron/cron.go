package cron

import (
	"log/slog"
	"os"
	"time"

	"git.my-itclub.ru/bots/school/internal/school"
	tg "git.my-itclub.ru/bots/school/internal/telegram"
	"github.com/robfig/cron/v3"
)

func RunTask() {
	l, _ := time.LoadLocation("Europe/Moscow")

	c := cron.New()
	cron.WithLocation(l)

	_, err := c.AddFunc(os.Getenv("SCHOOL_CRON"), func() {
		var messageError, message string

		grades, err := school.GetGrades(
			&school.Site{
				JWT:        os.Getenv("SCHOOL_JWT"),
				URL:        os.Getenv("SCHOOL_URL"),
				EucationID: os.Getenv("SCHOOL_EUCATION_ID"),
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

		slog.Info("Cron task completed")
	})
	if err != nil {
		slog.Warn("Error adding cron task", "error", err)
	}

	c.Start()
}