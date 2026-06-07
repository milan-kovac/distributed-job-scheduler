package router

import (
	"net/http"

	"github.com/milan-kovac/distributed-job-scheduler/api/internal/handler"
)

func New(apiHandler *handler.Handler) http.Handler {
	router := http.NewServeMux()

	// job
	router.HandleFunc("POST /jobs", apiHandler.CreateJob)
	router.HandleFunc("GET /jobs", apiHandler.ListJobs)
	router.HandleFunc("GET /jobs/{id}", apiHandler.GetJob)
	router.HandleFunc("DELETE /jobs/{id}", apiHandler.DeleteJob)
	router.HandleFunc("PATCH /jobs/{id}/cancel", apiHandler.CancelJob)

	return router
}
