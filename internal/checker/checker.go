package checker

import (
	"fmt"
	"os"
)

var (
	envVars = []string{
		"SCHOOL_HOST",
		"SCHOOL_LOGIN",
		"SCHOOL_PASSWORD",
		"SCHOOL_EUCATION_ID",
		"SCHOOL_TELEGRAM_TOKEN",
		"SCHOOL_CHAT_ID",
		"SCHOOL_CRON_WORK_WEEK",
		"SCHOOL_USER_AGENT",
	}
)

func CheckEnvVars() error {
	for _, v := range envVars {
		if os.Getenv(v) == "" {
			return fmt.Errorf("variable " + v + " is not set")
		}
	}

	return nil
}
