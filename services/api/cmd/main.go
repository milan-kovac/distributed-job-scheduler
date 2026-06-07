package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/milan-kovac/distributed-job-scheduler/api/internal/handler"
	"github.com/milan-kovac/distributed-job-scheduler/api/internal/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	apiHandler := handler.NewHandler(logger)
	apiRouter := router.New(apiHandler)

	logger.Info("API starting", "address", ":8080")

	if err := http.ListenAndServe(":8080", apiRouter); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}
