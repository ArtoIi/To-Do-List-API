package todoservice

import (
	"errors"
	"time"

	todoDTO "github.com/ArtoIi/To-Do-List-API/internal/application/todo_dto"
	"github.com/ArtoIi/To-Do-List-API/internal/domain"
	"github.com/ArtoIi/To-Do-List-API/internal/infrastructure/utils"
)

type ToDoService struct {
	repo domain.ToDoRepository
}

func NewToDoService(repo domain.ToDoRepository) *ToDoService {
	return &ToDoService{repo: repo}
}

func (s *ToDoService) CreatePost(dto *todoDTO.DTO, user_id int) (string, error) {
	if err := utils.ValidateStruct(dto); err != nil {
		return "", err
	}
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

func (s *ToDoService) GetByUserId(user_id int, filter todoDTO.Filter) ([]*domain.ToDo, int, error) {
	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	offset := (filter.Page - 1) * filter.Limit
	return s.repo.GetUserId(user_id, filter.Limit, offset)
}

func (s *ToDoService) DeletePost(id, userId int) error {
	todo, err := s.GetById(id)
	if err != nil {
		return err
	}

	if todo.UserID != userId {
		return errors.New("permissão negada")
	}
	return s.repo.Delete(id)
}

func (s *ToDoService) UpdatePost(dto *todoDTO.DTO, id int, userId int) (*domain.ToDo, error) {
	toDoNew, err := s.GetById(id)
	if err != nil {
		return nil, err
	}
	if toDoNew.UserID != userId {
		return nil, errors.New("permissão negada")
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
