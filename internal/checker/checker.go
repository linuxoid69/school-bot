package checker

import (
	"log/slog"
	"os"
)

var (
	envVars = []string{
		"SCHOOL_JWT",
		"SCHOOL_URL",
		"SCHOOL_EUCATION_ID",
		"SCHOOL_TOKEN",
		"SCHOOL_CHAT_ID",
		"SCHOOL_CRON",
	}
)

func CheckEnvVars() bool {
	for _, v := range envVars {
		if os.Getenv(v) == "" {
			slog.Error("variable " + v + " is not set")

			return false
		}
	}

	return true
}
