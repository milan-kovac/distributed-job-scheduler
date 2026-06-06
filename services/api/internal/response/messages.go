package response

var messages = map[Code]string{
	OK:        "ok",
	Created:   "created",
	Accepted:  "accepted",
	NoContent: "",

	BadRequest:          "bad request",
	Unauthorized:        "unauthorized",
	Forbidden:           "forbidden",
	NotFound:            "not found",
	Conflict:            "conflict",
	UnprocessableEntity: "unprocessable entity",
	TooManyRequests:     "too many requests",

	InternalError:      "internal server error",
	BadGateway:         "bad gateway",
	ServiceUnavailable: "service unavailable",
	GatewayTimeout:     "gateway timeout",
}

func (c Code) Message() string {
	if m, ok := messages[c]; ok {
		return m
	}
	return "unknown"
}
