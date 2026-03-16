package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	p_error "github.com/ArtoIi/To-Do-List-API/internal/infrastructure/error"
)

type Response[T any] struct {
	Data T               `json:"data"`
	Meta *PaginationMeta `json:"meta,omitempty"`
}

type PaginationMeta struct {
	Page  int `json:"current"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

func Respond[T any](w http.ResponseWriter, status int, data T, meta ...*PaginationMeta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := Response[T]{
		Data: data,
	}

	if len(meta) > 0 && meta[0] != nil {
		res.Meta = meta[0]

	}

	json.NewEncoder(w).Encode(res)
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondError(w http.ResponseWriter, r *http.Request, err error) {
	status := http.StatusInternalServerError
	message := "internal server error"

	var webErr p_error.Error
	if errors.As(err, &webErr) {
		status = webErr.Status
		message = webErr.Error()
	}

	log.Printf("[ERROR] %s %s: %v", r.Method, r.URL.Path, err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
