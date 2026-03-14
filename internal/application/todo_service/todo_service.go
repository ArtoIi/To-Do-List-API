package todoservice

import (
	"time"

	todoDTO "github.com/ArtoIi/To-Do-List-API/internal/application/todo_dto"
	"github.com/ArtoIi/To-Do-List-API/internal/domain"
)

type ToDoService struct {
	repo domain.ToDoRepository
}

func NewToDoService(repo domain.ToDoRepository) *ToDoService {
	return &ToDoService{repo: repo}
}

func (s *ToDoService) CreatePost(dto *todoDTO.CreateToDoDTO, user_id int) (string, error) {

	timenow := time.Now()
	todo := &domain.ToDo{
		UserID:      user_id,
		Title:       dto.Title,
		Description: dto.Description,
		CreatedAt:   timenow,
		UpdatedAt:   timenow,
	}

	if err := s.repo.Post(todo); err != nil {
		return "", err
	}
	return "Post Criado", nil

}
