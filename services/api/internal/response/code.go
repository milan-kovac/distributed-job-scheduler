package response

import "net/http"

// Code is an HTTP status code used by the API.
type Code int

const (
	// 2xx — success
	OK        Code = http.StatusOK        // 200 — success (GET, idempotent actions)
	Created   Code = http.StatusCreated   // 201 — resource created (POST with result)
	Accepted  Code = http.StatusAccepted  // 202 — accepted for async processing
	NoContent Code = http.StatusNoContent // 204 — success with no body (DELETE)

	// 4xx — client errors
	BadRequest          Code = http.StatusBadRequest          // 400 — bad JSON, syntax validation
	Unauthorized        Code = http.StatusUnauthorized        // 401 — missing/invalid token
	Forbidden           Code = http.StatusForbidden           // 403 — authenticated, not allowed
	NotFound            Code = http.StatusNotFound            // 404 — resource not found
	Conflict            Code = http.StatusConflict            // 409 — duplicate / state conflict
	UnprocessableEntity Code = http.StatusUnprocessableEntity // 422 — JSON valid, semantics not
	TooManyRequests     Code = http.StatusTooManyRequests     // 429 — rate limit

	// 5xx — server errors
	InternalError      Code = http.StatusInternalServerError // 500 — unexpected error
	BadGateway         Code = http.StatusBadGateway          // 502 — upstream (scheduler) not responding
	ServiceUnavailable Code = http.StatusServiceUnavailable  // 503 — service currently down
	GatewayTimeout     Code = http.StatusGatewayTimeout      // 504 — upstream timeout
)

// Int returns the numeric HTTP status code.
func (c Code) Int() int { return int(c) }
