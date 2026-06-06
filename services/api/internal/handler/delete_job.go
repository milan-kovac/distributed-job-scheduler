package handler

import (
	"net/http"

	"github.com/milan-kovac/distributed-job-scheduler/api/internal/response"
)

func (handler *Handler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	handler.log.Info("received DELETE /jobs/{id}", "id", id)
	response.Success(w, response.OK, nil)
}
