package todoHandler

import (
	"encoding/json"
	"net/http"

	todoDTO "github.com/ArtoIi/To-Do-List-API/internal/application/todo_dto"
	p_error "github.com/ArtoIi/To-Do-List-API/internal/infrastructure/error"
	"github.com/ArtoIi/To-Do-List-API/internal/infrastructure/utils"
)

type ToDoService interface {
	CreatePost(dto *todoDTO.CreateToDoDTO, user_id int) (string, error)
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

	var dto todoDTO.CreateToDoDTO
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
func (h *ToDoHandler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	utils.RespondError(w, r, err)
}
