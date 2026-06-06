package response

import (
	"encoding/json"
	"net/http"
)

type Body struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Success(w http.ResponseWriter, code Code, data any) {
	if code == NoContent {
		w.WriteHeader(code.Int())
		return
	}
	write(w, code, code.Message(), data)
}

func Error(w http.ResponseWriter, code Code) {
	write(w, code, code.Message(), nil)
}

func ErrorWith(w http.ResponseWriter, code Code, msg string) {
	write(w, code, msg, nil)
}

func write(w http.ResponseWriter, code Code, msg string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code.Int())
	_ = json.NewEncoder(w).Encode(Body{
		Code:    code.Int(),
		Message: msg,
		Data:    data,
	})
}
