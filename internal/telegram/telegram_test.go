package telegram

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/linuxoid69/school-bot/internal/school"
)

// func TestMessage_SendGrades(t *testing.T) {
// 	type fields struct {
// 		Text   string
// 		ChatID string
// 		Token  string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			m := &Message{
// 				Text:   tt.fields.Text,
// 				ChatID: tt.fields.ChatID,
// 				Token:  tt.fields.Token,
// 			}
// 			if err := m.SendGrades(); (err != nil) != tt.wantErr {
// 				t.Errorf("Message.SendGrades() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

func TestCreateMessage(t *testing.T) {
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
			got, err := CreateMessage(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("CreateMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
