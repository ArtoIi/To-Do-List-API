package application

import (
	userDTO "github.com/ArtoIi/To-Do-List-API/internal/application/user_dto"
	"github.com/ArtoIi/To-Do-List-API/internal/domain"
)

type ToDoService struct {
	repo domain.ToDoRepository
}

func NewToDoService(repository domain.ToDoRepository) *ToDoService {
	return &ToDoService{repo: repository}
}

func (s *ToDoService) CreateUser(dto userDTO.CreateUserDTO) (string, error) {
	err := s.repo.Register(dto)

	token, err := dto.GenerateToken()

	if err != nil {
		return "", nil
	}

	return token, err
}

func (s *ToDoService) GetByEmail(email string) (*domain.User, error) {
	user, err := s.repo.GetEmail(email)

	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *ToDoService) GetById(id int) (*domain.User, error) {
	user, err := s.repo.GetId(id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *ToDoService) DeleteUser(id int) error {
	err := s.repo.Delete(id)
	return err
}

func (s *ToDoService) UpdateUser(dto domain.CreateUserDTO, id int) error {
	err := s.repo.Update(dto, id)

	return err
}
