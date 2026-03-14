package todoHandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	todoDTO "github.com/ArtoIi/To-Do-List-API/internal/application/todo_dto"
	"github.com/ArtoIi/To-Do-List-API/internal/domain"
	p_error "github.com/ArtoIi/To-Do-List-API/internal/infrastructure/error"
	"github.com/ArtoIi/To-Do-List-API/internal/infrastructure/utils"
)

type ToDoService interface {
	CreatePost(dto *todoDTO.ToDoDTO, user_id int) (string, error)
	GetById(id int) (*domain.ToDo, error)
	GetByUserId(user_id int) ([]*domain.ToDo, error)
	DeletePost(id int) error
	UpdatePost(dto *todoDTO.ToDoDTO, id int) (*domain.ToDo, error)
}

type ToDoHandler struct {
	service ToDoService
}

func NewToDoHandler(service ToDoService) *ToDoHandler {
	return &ToDoHandler{service: service}
}

func (h ToDoHandler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.handleError(w, r, p_error.ErrInvalidMethod)
		return
	}

	var dto todoDTO.ToDoDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		h.handleError(w, r, p_error.ErrInvalidJSON)
		return
	}

	idVal := r.Context().Value("user_id")

	userIDFloat, ok := idVal.(float64)
	if !ok {
		http.Error(w, "user_id inválido", http.StatusUnauthorized)
		return
	}

	userID := int(userIDFloat)

	resposta, err := h.service.CreatePost(&dto, userID)
	if err != nil {
		h.handleError(w, r, err)
		return
	}
	utils.Respond(w, http.StatusCreated, resposta)
}

func (h *ToDoHandler) GetId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.handleError(w, r, p_error.ErrInvalidMethod)
		return
	}
	idstr := r.PathValue("id")
	id, _ := strconv.Atoi(idstr)

	todo, err := h.service.GetById(id)
	if err != nil {
		h.handleError(w, r, err)
		return
	}
	utils.Respond(w, http.StatusOK, todo)

}
func (h *ToDoHandler) GetUserId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.handleError(w, r, p_error.ErrInvalidMethod)
		return
	}
	idstr := r.PathValue("user_id")
	id, _ := strconv.Atoi(idstr)

	todo, err := h.service.GetByUserId(id)
	if err != nil {
		h.handleError(w, r, err)
		return
	}
	utils.Respond(w, http.StatusOK, todo)

}

func (h *ToDoHandler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	utils.RespondError(w, r, err)
}
