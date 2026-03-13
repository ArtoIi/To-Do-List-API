package p_error

import (
	"errors"
	"net/http"

	"github.com/ArtoIi/To-Do-List-API/internal/infrastructure/utils"
)

var (
	ErrInvalidMethod  = utils.Error{Err: errors.New("Invalid method"), Status: http.StatusBadRequest}
	ErrInvalidJSON    = utils.Error{Err: errors.New("Invalid JSON"), Status: http.StatusBadRequest}
	ErrInvalidID      = utils.Error{Err: errors.New("author id is required"), Status: http.StatusBadRequest}
	ErrInvalidAccount = utils.Error{Err: errors.New("Invalid email or password"), Status: http.StatusUnauthorized}
	ErrNotFound       = utils.Error{Err: errors.New("event not found"), Status: http.StatusNotFound}
)
