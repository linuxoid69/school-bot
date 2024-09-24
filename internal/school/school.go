package school

import (
	"context"
	"fmt"
	"io"
	"net/http"
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

func GetGrades(site *Site) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, site.URL, nil)
	if err != nil {
		return nil, err
	}

	todayDate := time.Now().Format("02.01.2006")

	if site.DateFrom == "" && site.DateTo == "" {
		site.DateFrom = todayDate
		site.DateTo = todayDate
	}

	req.Header.Set("Cookie", fmt.Sprintf("X-JWT-Token=%s", site.JWT))

	q := req.URL.Query()
	q.Add("p_educations[]", site.EucationID)
	q.Add("p_date_from", site.DateFrom)
	q.Add("p_date_to", site.DateTo)
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
