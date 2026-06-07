package handler

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	log      *slog.Logger
	validate *validator.Validate
}

func NewHandler(logger *slog.Logger) *Handler {
	validate := validator.New()

	return &Handler{
		log:      logger,
		validate: validate,
	}
}
