package userHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	userDTO "github.com/ArtoIi/To-Do-List-API/internal/application/user_dto"
	"github.com/ArtoIi/To-Do-List-API/internal/domain"
	p_error "github.com/ArtoIi/To-Do-List-API/internal/infrastructure/error"
	"github.com/ArtoIi/To-Do-List-API/internal/infrastructure/utils"
)

type UserService interface {
	CreateUser(dto userDTO.CreateUserDTO) (string, error)
	GetByEmail(email string) (*domain.User, error)
	GetById(id int) (*domain.User, error)
	DeleteUser(id int) error
	UpdateUser(dto userDTO.UpdateUserDTO) (*domain.User, error)
	Login(email, password string) (string, error)
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(s UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.handleError(w, r, p_error.ErrInvalidMethod)
		return
	}
	var user userDTO.CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.handleError(w, r, p_error.ErrInvalidJSON)
		return
	}

	token, err := h.service.CreateUser(user)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusCreated, token)
}

func (h *UserHandler) GetEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.handleError(w, r, p_error.ErrInvalidMethod)
		return
	}
	email := r.PathValue("email")

	result, err := h.service.GetByEmail(email)
	if err != nil {
		h.handleError(w, r, p_error.ErrInvalidID)
	}

	utils.Respond(w, http.StatusOK, result)

}
func (h *UserHandler) GetId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.handleError(w, r, p_error.ErrInvalidMethod)
		return
	}

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	result, err := h.service.GetById(id)
	if err != nil {
		h.handleError(w, r, p_error.ErrInvalidID)
	}

	utils.Respond(w, http.StatusOK, result)

}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.handleError(w, r, p_error.ErrInvalidMethod)
		return
	}

	var dto userDTO.UpdateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		h.handleError(w, r, err)
		return
	}

	updatedUser, err := h.service.UpdateUser(dto)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, updatedUser)

}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.handleError(w, r, p_error.ErrInvalidMethod)
		return
	}

	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)

	err := h.service.DeleteUser(id)
	if err != nil {
		h.handleError(w, r, err)
	}
	utils.Respond(w, http.StatusNoContent, "")

}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.handleError(w, r, p_error.ErrInvalidMethod)
		return
	}

	var dto userDTO.LoginUserDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		h.handleError(w, r, err)
		return
	}
	token, err := h.service.Login(dto.Email, dto.Password)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, token)

}

// Brincadeirinha
func (h *UserHandler) Identify(w http.ResponseWriter, r *http.Request) {
	idVal := r.Context().Value("user_id")

	userIDFloat, ok := idVal.(float64)
	if !ok {
		http.Error(w, "user_id inválido", http.StatusUnauthorized)
		return
	}

	userID := int(userIDFloat)

	user, err := h.service.GetById(userID)
	if err != nil {
		h.handleError(w, r, err)
	}
	frase := fmt.Sprintf("Oii, %s", user.Name)
	utils.Respond(w, http.StatusOK, frase)
}

func (h *UserHandler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	utils.RespondError(w, r, err)
}
