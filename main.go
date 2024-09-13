package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
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

	CreateMessage(data)
}

func GetGrades(site *Site) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s", site.URL), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Cookie", fmt.Sprintf("X-JWT-Token=%s", site.JWT))

	todayDate := time.Now().Format("02.01.2006")

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

func CreateMessage(data []byte) {
	var grades Grades

	json.Unmarshal(data, &grades)

	for _, item := range grades.Data.Items {
		if item.EstimateComment == nil {
			item.EstimateComment = ""
		}

		fmt.Printf("\nУрок: %s \nДата: %s \nИтог: %s \nГде: %s \nКомментарий: %s \n----------------------------------",
			item.SubjectName,
			item.Date,
			item.EstimateValueName,
			item.EstimateTypeName,
			item.EstimateComment)
	}
}

func SendGrades() {

}
