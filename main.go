package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

type Grades struct {
	Data Data `json:"data"`
}

type Items struct {
	SubjectName       string `json:"subject_name"`
	Date              string `json:"date"`
	EstimateValueName string `json:"estimate_value_name"`
	EstimateTypeName  string `json:"estimate_type_name"`
	EstimateComment   any    `json:"estimate_comment"`
}
type Data struct {
	Items []Items `json:"items"`
}

type Site struct {
	JWT        string
	URL        string
	EucationID string
	DateFrom   string
	DateTo     string
}

type Message struct {
	Text   string `json:"text"`
	ChatID string `json:"chat_id"`
	Token  string
}

func main() {
	data, err := GetGrades(
		&Site{
			JWT:        os.Getenv("JWT"),
			URL:        os.Getenv("URL"),
			EucationID: os.Getenv("EUCATION_ID"),
			DateFrom:   os.Getenv("DATE_FROM"),
			DateTo:     os.Getenv("DATE_TO")},
	)
	if err != nil {
		slog.Warn("Error getting grades", "error", err)
	}

	messageString, err := CreateMessage(data)
	if err != nil {
		slog.Warn("Error creating message", "error", err)
	}

	mesg := Message{
		Text:   messageString,
		ChatID: os.Getenv("CHAT_ID"),
		Token:  os.Getenv("TOKEN"),
	}

	err = mesg.SendGrades()
	if err != nil {
		slog.Warn("Error sending message", "error", err)
	}
}

func GetGrades(site *Site) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s", site.URL), nil)
	if err != nil {
		return nil, err
	}

	todayDate := time.Now().Format("02.01.2006")

	req.Header.Set("Cookie", fmt.Sprintf("X-JWT-Token=%s", site.JWT))

	q := req.URL.Query()
	q.Add("p_educations[]", fmt.Sprintf("%s", site.EucationID))
	q.Add("p_date_from", todayDate)
	q.Add("p_date_to", todayDate)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func CreateMessage(data []byte) (string, error) {
	var grades Grades

	json.Unmarshal(data, &grades)

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(
		"Оценки за %s: \n%s\n",
		time.Now().Format("02.01.2006"),
		"========================================"))

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
