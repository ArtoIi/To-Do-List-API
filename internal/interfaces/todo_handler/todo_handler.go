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
	CreatePost(dto *todoDTO.DTO, user_id int) (string, error)
	GetById(id int) (*domain.ToDo, error)
	GetByUserId(user_id int, filter todoDTO.Filter) ([]*domain.ToDo, int, error)
	DeletePost(id, userId int) error
	UpdatePost(dto *todoDTO.DTO, id int, userId int) (*domain.ToDo, error)
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

	var dto todoDTO.DTO
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
	userIDstr := r.PathValue("user_id")
	userID, _ := strconv.Atoi(userIDstr)

	query := r.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	if page <= 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit <= 0 {
		limit = 10
	}

	filter := todoDTO.Filter{
		Page:  page,
		Limit: limit,
	}

	todos, total, err := h.service.GetByUserId(userID, filter)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	meta := &utils.PaginationMeta{
		Page:  filter.Page,
		Limit: limit,
		Total: total,
	}

	utils.Respond(w, http.StatusOK, todos, meta)
}
func (h *ToDoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.handleError(w, r, p_error.ErrInvalidMethod)
		return
	}

	userid := r.Context().Value("user_id").(float64)

	idstr := r.PathValue("id")
	id, _ := strconv.Atoi(idstr)

	if err := h.service.DeletePost(id, int(userid)); err != nil {
		h.handleError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusNoContent, "")

}
func (h *ToDoHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.handleError(w, r, p_error.ErrInvalidMethod)
		return
	}

	userid := r.Context().Value("user_id").(float64)

	idstr := r.PathValue("id")
	id, _ := strconv.Atoi(idstr)

	var dto *todoDTO.DTO
	json.NewDecoder(r.Body).Decode(&dto)

	newToDo, err := h.service.UpdatePost(dto, id, int(userid))
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	utils.Respond(w, http.StatusOK, newToDo)

}

func (h *ToDoHandler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	utils.RespondError(w, r, err)
}
