package telegram

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/linuxoid69/school-bot/internal/school"
)

func TestCreateTodayReport(t *testing.T) {
	var grades, gradesEmpty school.Grades

	today := time.Now().Format("02.01.2006")

	grades.Data.Items = []school.Items{
		{
			SubjectName:       "Физическая культура",
			Date:              today,
			EstimateValueName: "3",
			EstimateTypeName:  "Работа на уроке",
			EstimateComment:   nil,
		},
	}

	gradesEmpty.Data.Items = []school.Items{}

	data, err := json.Marshal(grades)
	if err != nil {
		fmt.Errorf("Error marshal json")
	}

	dataEmpty, err := json.Marshal(gradesEmpty)
	if err != nil {
		fmt.Errorf("Error marshal json")
	}

	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Create message",
			args: args{
				data: []byte(data),
			},
			want: fmt.Sprintf(`Оценки за %s:
========================================

Урок: Физическая культура
Итог: 3
Где: Работа на уроке
Комментарий:
-------------------------------------------------------`, today),
			wantErr: false,
		},
		{
			name: "Create empty message",
			args: args{
				data: []byte(""),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Data is empty",
			args: args{
				data: []byte(dataEmpty),
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateTodayReport(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTodayReport() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if got != tt.want {
				t.Errorf("CreateTodayReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
