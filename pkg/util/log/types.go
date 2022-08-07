package log

import (
	"time"
)

type LogError struct {
	ID           string    `json:"id"`
	Header       string    `json:"request_header"`
	Body         string    `json:"request_body"`
	URL          string    `json:"url"`
	HttpMethod   string    `json:"http_method"`
	Email        string    `json:"email"`
	ErrorMessage string    `json:"error_message"`
	Level        string    `json:"level"`
	AppName      string    `json:"app_name"`
	Version      string    `json:"version"`
	Env          string    `json:"env"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
