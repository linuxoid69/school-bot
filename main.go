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

func main() {
	var dateFrom, dateTo string

	flag.StringVar(&dateFrom, "from", "", "use as 01.01.2001")
	flag.StringVar(&dateTo, "to", "", "use as 01.01.2001")
	flag.Parse()

	if dateFrom != "" && dateTo != "" {
		grades, err := school.GetGrades(
			&school.Site{
				JWT:        os.Getenv("SCHOOL_JWT"),
				URL:        os.Getenv("SCHOOL_URL"),
				EucationID: os.Getenv("SCHOOL_EUCATION_ID"),
				DateFrom:   dateFrom,
				DateTo:     dateTo,
			},
		)
		if err != nil {
			slog.Warn("Error getting grades", "error", err)
		}

		fmt.Println(string(grades))
	} else {
		slog.Info("Start bot school")

		checker.CheckEnvVars()

		cron.RunTask()

		for {
			time.Sleep(time.Millisecond * 500)
		}
	}
}
