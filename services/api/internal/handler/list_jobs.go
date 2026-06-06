package handler

import (
	"net/http"

	"github.com/milan-kovac/distributed-job-scheduler/api/internal/response"
)

func (handler *Handler) ListJobs(w http.ResponseWriter, r *http.Request) {
	handler.log.Info("received GET /jobs")

	response.Success(w, response.OK, nil)
}
