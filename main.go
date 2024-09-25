package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"git.my-itclub.ru/bots/school/internal/checker"
	"git.my-itclub.ru/bots/school/internal/cron"
	"git.my-itclub.ru/bots/school/internal/school"
)

var Version string

func main() {
	var (
		dateFromFlag, dateToFlag string
		versionFlag              bool
	)

	flag.StringVar(&dateFromFlag, "from", "", "use as 01.01.2001")
	flag.StringVar(&dateToFlag, "to", "", "use as 01.01.2001")
	flag.BoolVar(&versionFlag, "v", false, "show version")
	flag.Parse()

	if versionFlag {
		fmt.Println(Version)

		return
	}

	if dateFromFlag != "" && dateToFlag != "" {
		grades, err := school.GetGrades(
			&school.Site{
				JWT:        os.Getenv("SCHOOL_JWT"),
				URL:        os.Getenv("SCHOOL_URL"),
				EucationID: os.Getenv("SCHOOL_EUCATION_ID"),
				UserAgent:  os.Getenv("SCHOOL_USER_AGENT"),
				DateFrom:   dateFromFlag,
				DateTo:     dateToFlag,
			},
		)
		if err != nil {
			slog.Warn("Error getting grades", "error", err)
		}

		fmt.Println(string(grades))
	} else {
		slog.Info("Start bot school")

		if err := checker.CheckEnvVars(); err != nil {
			slog.Error("Error checking env vars", "error", err)

			return
		}

		cron.RunTask()

		for {
			time.Sleep(time.Millisecond * 500)
		}
	}
}
