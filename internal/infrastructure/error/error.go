package p_error

import (
	"errors"
	"net/http"
)

type Error struct {
	Err    error
	Status int
}

func (e Error) Error() string {
	return e.Err.Error()
}

var (
	ErrInvalidMethod  = Error{Err: errors.New("Invalid method"), Status: http.StatusBadRequest}
	ErrInvalidJSON    = Error{Err: errors.New("Invalid JSON"), Status: http.StatusBadRequest}
	ErrInvalidID      = Error{Err: errors.New("author id is required"), Status: http.StatusBadRequest}
	ErrInvalidAccount = Error{Err: errors.New("Invalid email or password"), Status: http.StatusUnauthorized}
	ErrNotFound       = Error{Err: errors.New("event not found"), Status: http.StatusNotFound}
)
