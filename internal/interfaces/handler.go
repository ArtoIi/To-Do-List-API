package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/ArtoIi/To-Do-List-API/internal/application"
	"github.com/ArtoIi/To-Do-List-API/internal/domain"
)

type ToDoHandler struct {
	service *application.ToDoService
}

func NewToDoHandler(s *application.ToDoService) *ToDoHandler {
	return &ToDoHandler{service: s}
}

func (h *ToDoHandler) Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Metodo invalido", http.StatusBadRequest)
		return
	}
	var user domain.CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "JSON Invalido", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Registrado com Sucesso!")
}
