package handler

import (
	"net/http"

	"github.com/milan-kovac/distributed-job-scheduler/api/internal/model"
	"github.com/milan-kovac/distributed-job-scheduler/api/internal/response"
)

func (handler *Handler) CreateJob(w http.ResponseWriter, r *http.Request) {
	var req model.CreateJobRequest

	if err := handler.BindAndValidate(r, &req); err != nil {
		response.ErrorWith(w, response.BadRequest, err.Error())
		return
	}

	handler.log.Info("create job request validated",
		"type", req.Type,
		"url", req.Payload.URL,
		"method", req.Payload.Method,
		"delay_ms", req.DelayMS,
		"cron_expr", req.CronExpr,
		"run_at", req.RunAt,
		"max_attempts", req.MaxAttempts,
	)

	response.Success(w, response.OK, nil)
}
