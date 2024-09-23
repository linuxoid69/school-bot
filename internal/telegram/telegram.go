package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"unicode"

	"git.my-itclub.ru/bots/school/internal/school"
)

type Message struct {
	Text   string `json:"text"`
	ChatID string `json:"chat_id"`
	Token  string
}

func (m *Message) SendGrades() error {
	payload, err := json.Marshal(m)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", m.Token),
		bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func CreateMessage(data []byte) (string, error) {
	var grades school.Grades

	if err := json.Unmarshal(data, &grades); err != nil {
		return "", err
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(
		"Оценки за %s: \n%s\n",
		time.Now().Format("02.01.2006"),
		"========================================"))

	if len(grades.Data.Items) == 0 {
		return "", nil
	}

	for _, item := range grades.Data.Items {
		if item.EstimateComment == nil {
			item.EstimateComment = ""
		}

		_, err := sb.WriteString(
			fmt.Sprintf("\nУрок: %s \nИтог: %s \nГде: %s \nКомментарий: %s \n%s",
				item.SubjectName,
				item.EstimateValueName,
				item.EstimateTypeName,
				item.EstimateComment,
				"-------------------------------------------------------"))
		if err != nil {
			return "", err
		}
	}

	return sb.String(), nil
}

func CreateWeekReport(fromDate, toDate string, data []byte) (string, error) {
	var grades school.Grades

	if err := json.Unmarshal(data, &grades); err != nil {
		return "", err
	}

	weekGrades := make(map[string]string)
	for _, item := range grades.Data.Items {
		if unicode.IsDigit([]rune(item.EstimateValueName)[0]) {
			weekGrades[item.SubjectName] += item.EstimateValueName + ","
		}
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(
		"Отчет за неделю: %s-%s\n%s\n", fromDate, toDate,
		"========================================"))

	for subject, grade := range weekGrades {
		sb.WriteString(fmt.Sprintf("%s: %s\n", subject, grade))
	}

	return sb.String(), nil
}
