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

func (s *ToDoService) CreatePost(dto *todoDTO.ToDoDTO, user_id int) (string, error) {

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

func (s *ToDoService) GetById(id int) (*domain.ToDo, error) {
	return s.repo.GetId(id)
}

func (s *ToDoService) GetByUserId(user_id int) ([]*domain.ToDo, error) {
	return s.repo.GetUserId(user_id)
}

func (s *ToDoService) DeletePost(id int) error {
	return s.repo.Delete(id)
}

func (s *ToDoService) UpdatePost(dto *todoDTO.ToDoDTO, id int) (*domain.ToDo, error) {
	toDoNew, err := s.GetById(id)
	if err != nil {
		return nil, err
	}
	if dto.Title != "" {
		toDoNew.Title = dto.Title
	}
	if dto.Description != "" {
		toDoNew.Description = dto.Description
	}
	toDoNew.UpdatedAt = time.Now()

	_, err = s.repo.Update(toDoNew)
	if err != nil {
		return nil, err
	}
	return toDoNew, nil
}
