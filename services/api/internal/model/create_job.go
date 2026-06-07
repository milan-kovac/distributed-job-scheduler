package model

import "time"

// CreateJobRequest POST /jobs.

type CreateJobRequest struct {
	Type JobType `json:"type" validate:"required,oneof=once delayed cron"`

	RunAt    *time.Time `json:"run_at,omitempty" validate:"required_if=Type once"`
	DelayMS  int64      `json:"delay_ms,omitempty" validate:"required_if=Type delayed,gte=0"`
	CronExpr string     `json:"cron_expr,omitempty" validate:"required_if=Type cron"`

	Payload HTTPPayload `json:"payload" validate:"required"`

	MaxAttempts int `json:"max_attempts,omitempty" validate:"gte=0"`
}
