package handler

import (
	"io"
	"net/http"

	"github.com/milan-kovac/distributed-job-scheduler/api/internal/response"
)

func (handler *Handler) CreateJob(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, response.BadRequest)
		return
	}

	handler.log.Info("received POST /jobs", "body", string(body))
	response.Success(w, response.OK, nil)
}
