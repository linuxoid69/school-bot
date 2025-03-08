package cron

import (
	"log/slog"
	"os"

	"github.com/linuxoid69/school-bot/internal/school"
	tg "github.com/linuxoid69/school-bot/internal/telegram"
)

type Report struct {
	Type string
}

func (r *Report) BuildReport(token string, site *school.Site) {
	var messageError, message string

	grades, err := site.GetGrades()
	if err != nil {
		slog.Warn("Error getting grades", "error", err)

		messageError = school.ERROR_GETTIG_GRADES
	}

	switch r.Type {
	case "week":
		message, err = tg.CreateWeekReport(site.DateFrom, site.DateTo, grades)
		if err != nil {
			slog.Warn("Error creating message", "error", err)

			messageError = school.ERROR_CREATING_MESSAGE
		}
	case "today":
		message, err = tg.CreateTodayReport(grades)
		if err != nil {
			slog.Warn("Error creating message", "error", err)

			messageError = school.ERROR_CREATING_MESSAGE
		}
	}

	if messageError != "" {
		message = messageError
	}

	mesg := tg.Message{
		Text:   message,
		ChatID: os.Getenv("SCHOOL_CHAT_ID"),
		Token:  os.Getenv("SCHOOL_TELEGRAM_TOKEN"),
	}

	if err = mesg.SendGrades(); err != nil {
		slog.Warn("Error sending message", "error", err)
	}

	slog.Info("Cron task report completed", "type", r.Type)
}
