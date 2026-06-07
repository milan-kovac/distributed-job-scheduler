package handler

import (
	"net/http"

	"github.com/milan-kovac/distributed-job-scheduler/api/internal/response"
)

func (handler *Handler) CancelJob(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	handler.log.Info("received POST /jobs/{id}/cancel", "id", id)

	response.Success(w, response.OK, nil)
}
