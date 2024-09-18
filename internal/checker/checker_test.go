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
		want    bool
	}{
		{
			name: "CheckEnvVars",
			want: true,
			envVars: map[string]string{
				"SCHOOL_JWT":            "123",
				"SCHOOL_URL":            "https://example.com",
				"SCHOOL_CHAT_ID":        "1234",
				"SCHOOL_EUCATION_ID":    "1234",
				"SCHOOL_TOKEN":          "1234",
				"SCHOOL_CRON_WORK_WEEK": "* * * * *",
			},
		},
		{
			name: "CheckEnvVars not set",
			want: false,
			envVars: map[string]string{
				"SCHOOL_JWT":            "",
				"SCHOOL_URL":            "",
				"SCHOOL_CHAT_ID":        "",
				"SCHOOL_EUCATION_ID":    "",
				"SCHOOL_TOKEN":          "",
				"SCHOOL_CRON_WORK_WEEK": "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EnvVars(t, tt.envVars)

			if got := CheckEnvVars(); got != tt.want {
				t.Errorf("CheckEnvVars() = %v, want %v", got, tt.want)
			}
		})
	}
}
