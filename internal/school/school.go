package school

import (
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
