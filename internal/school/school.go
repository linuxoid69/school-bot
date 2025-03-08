package school

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	LOGIN_LOCATION         string = "api/user/auth/login"
	DEFAULT_LOGIN_TYPE     string = "email"
	JOURNAL_URL            string = "api/journal/estimate/table"
	ERROR_GETTIG_GRADES    string = "Ошибка получения оценок"
	ERROR_CREATING_MESSAGE string = "Ошибка создания сообщения"
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
	JWT             string
	Host            string
	JournalLocation string
	EucationID      string
	DateFrom        string
	DateTo          string
	UserAgent       string
	Login           Login
}

type Login struct {
	Host     string `json:"host"`
	Type     string `json:"type"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
	Validations []any `json:"validations"`
	Messages    []any `json:"messages"`
	Debug       []any `json:"debug"`
}

func (login *Login) NewLogin() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	bodyJson, err := json.Marshal(login)
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, login.Host+"/"+LOGIN_LOCATION, bytes.NewBuffer(bodyJson))
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}

	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error: %v", err)
	}

	return string(resBody), nil
}

func (s *Site) GetGrades() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.JournalLocation, nil)
	if err != nil {
		return nil, err
	}

	todayDate := time.Now().Format("02.01.2006")

	if s.DateFrom == "" && s.DateTo == "" {
		s.DateFrom = todayDate
		s.DateTo = todayDate
	}

	req.Header.Set("Cookie", fmt.Sprintf("X-JWT-Token=%s", s.JWT))

	q := req.URL.Query()
	q.Add("p_educations[]", s.EucationID)
	q.Add("p_date_from", s.DateFrom)
	q.Add("p_date_to", s.DateTo)

	req.URL.RawQuery = q.Encode()
	req.Header.Set("User-Agent", s.UserAgent)

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

func (l *Login) GetJWTToken() (string, error) {
	loginResponse, err := l.NewLogin()
	if err != nil {
		return "", fmt.Errorf("Can't login %v", err)
	}

	var jwt *LoginResponse

	if err := json.Unmarshal([]byte(loginResponse), &jwt); err != nil {
		return "", fmt.Errorf("Can't unmarshal loginResponse %v", err)
	}

	return jwt.Data.Token, nil
}
