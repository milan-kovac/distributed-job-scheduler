package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/milan-kovac/distributed-job-scheduler/api/internal/handler"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	apiHandler := handler.NewHandler(logger)

	router := http.NewServeMux()
	router.HandleFunc("POST /jobs", apiHandler.CreateJob)
	router.HandleFunc("GET /jobs", apiHandler.ListJobs)
	router.HandleFunc("GET /jobs/{id}", apiHandler.GetJob)
	router.HandleFunc("DELETE /jobs/{id}", apiHandler.DeleteJob)

	logger.Info("API starting", "addr", ":8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
