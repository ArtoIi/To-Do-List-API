package application

import "github.com/ArtoIi/To-Do-List-API/internal/domain"

type ToDoService struct {
	repo domain.ToDoRepository
}

func NewToDoService(repository domain.ToDoRepository) *ToDoService {
	return &ToDoService{repo: repository}
}

func (s *ToDoService) CreateUser(dto domain.CreateUserDTO) error {
	err := s.repo.Register(dto)
	return err
}
