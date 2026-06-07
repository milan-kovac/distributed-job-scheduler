package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (h *Handler) BindAndValidate(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		return formatJSONError(err)
	}

	if err := h.validate.Struct(dst); err != nil {
		return formatValidationError(err)
	}

	return nil
}

func formatJSONError(err error) error {
	// unknown field
	if strings.Contains(err.Error(), "unknown field") {
		return fmt.Errorf("unknown field: %s", extractUnknownField(err.Error()))
	}

	var typeErr *json.UnmarshalTypeError
	if errors.As(err, &typeErr) {
		return fmt.Errorf("field '%s' has invalid type", strings.ToLower(typeErr.Field))
	}

	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		return fmt.Errorf("invalid JSON syntax")
	}

	return fmt.Errorf("invalid JSON")
}

func formatValidationError(err error) error {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return err
	}

	e := ve[0]

	// use JSON field name instead of Go struct name
	field := strings.ToLower(e.Field())

	switch e.Tag() {
	case "required":
		return fmt.Errorf("%s is required", field)

	case "required_if":
		return fmt.Errorf("%s is required", field)

	case "oneof":
		return fmt.Errorf("%s must be one of: %s", field, e.Param())

	case "duration":
		return fmt.Errorf("%s must be a valid duration (e.g. 10s, 5m, 1h)", field)

	case "gte":
		return fmt.Errorf("%s must be greater than or equal to %s", field, e.Param())

	default:
		return fmt.Errorf("%s is invalid", field)
	}
}

func extractUnknownField(msg string) string {
	parts := strings.Split(msg, "\"")
	if len(parts) >= 2 {
		return parts[1]
	}
	return "unknown"
}
