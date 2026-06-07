package model

type JobType string

const (
	JobTypeOnce    JobType = "once"
	JobTypeDelayed JobType = "delayed"
	JobTypeCron    JobType = "cron"
)

type HTTPPayload struct {
	URL     string      `json:"url" validate:"required,url"`
	Method  string      `json:"method" validate:"required,oneof=GET POST PUT PATCH DELETE"`
	Headers HTTPHeaders `json:"headers,omitempty"`
	Body    string      `json:"body,omitempty"`
}

type HTTPHeaders struct {
	Authorization string `json:"Authorization" validate:"required"`
	ContentType   string `json:"Content-Type" validate:"required"`
}
