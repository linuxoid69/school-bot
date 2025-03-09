package checker

import (
	"testing"
)

func EnvVars(t *testing.T, envVars map[string]string) {
	for k, v := range envVars {
		t.Setenv(k, v)
	}
}

func TestCheckEnvVars(t *testing.T) {
	tests := []struct {
		name    string
		envVars map[string]string
		wantErr bool
	}{
		{
			name:    "CheckEnvVars",
			wantErr: false,
			envVars: map[string]string{
				"SCHOOL_HOST":             "https://example.com",
				"SCHOOL_LOGIN":            "login",
				"SCHOOL_PASSWORD":         "secret",
				"SCHOOL_CHAT_ID":          "1234",
				"SCHOOL_EUCATION_ID":      "1234",
				"SCHOOL_TELEGRAM_TOKEN":   "1234",
				"SCHOOL_CRON_WORK_WEEK":   "* * * * *",
				"SCHOOL_CRON_WEEK_REPORT": "* * * * *",
				"SCHOOL_USER_AGENT":       "Mozilla/5.0 (X11; Linux x86_64)",
			},
		},
		{
			name:    "CheckEnvVars not set",
			wantErr: true,
			envVars: map[string]string{
				"SCHOOL_HOST":             "",
				"SCHOOL_LOGIN":            "",
				"SCHOOL_PASSWORD":         "",
				"SCHOOL_CHAT_ID":          "",
				"SCHOOL_EUCATION_ID":      "",
				"SCHOOL_TELEGRAM_TOKEN":   "",
				"SCHOOL_CRON_WORK_WEEK":   "",
				"SCHOOL_CRON_WEEK_REPORT": "",
				"SCHOOL_USER_AGENT":       "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EnvVars(t, tt.envVars)

			err := CheckEnvVars()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTodayReport() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
		})
	}
}

// SCHOOL_HOST            # "https://dnevnik2.petersburgedu.ru"  Хост сайта школы
// SCHOOL_LOGIN           # логин
// SCHOOL_PASSWORD        # пароль
// SCHOOL_EUCATION_ID     # id учащегося
// SCHOOL_TELEGRAM_TOKEN  # телеграм токен
// SCHOOL_CHAT_ID         # id чата телеграм
// SCHOOL_CRON_WORK_WEEK  # cron выражение когда будут опрашиваться данные со школы
// SCHOOL_USER_AGENT      # user-agent для запросов к сайту школы
