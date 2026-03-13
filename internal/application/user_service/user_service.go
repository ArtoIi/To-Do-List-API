package userService

import (
	"errors"
	"fmt"
	"strconv"

	userDTO "github.com/ArtoIi/To-Do-List-API/internal/application/user_dto"
	"github.com/ArtoIi/To-Do-List-API/internal/domain"
	p_error "github.com/ArtoIi/To-Do-List-API/internal/infrastructure/error"
	"github.com/ArtoIi/To-Do-List-API/internal/infrastructure/security"
)

type ToDoService struct {
	repo domain.UserRepository
}

func NewToDoService(repository domain.UserRepository) *ToDoService {
	return &ToDoService{repo: repository}
}

func (s *ToDoService) CreateUser(dto userDTO.CreateUserDTO) (string, error) {

	hashed, _ := security.HashedPassword(dto.Password)
	user := &domain.User{
		Name:           dto.Name,
		Email:          dto.Email,
		HashedPassword: hashed,
	}
	err := s.repo.Register(user)
	if err != nil {
		return "", err
	}

	token, err := security.GenerateToken(user)

	if err != nil {
		return "", err
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

func (s *ToDoService) UpdateUser(dto userDTO.UpdateUserDTO) (*domain.User, error) {
	if dto.Id == "" {
		return nil, errors.New("id is required")
	}

	id, err := strconv.Atoi(dto.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}

	oldUser, err := s.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	user := oldUser

	if dto.Name != nil {
		user.Name = *dto.Name
	}

	if dto.Email != nil {
		user.Email = *dto.Email
	}

	if dto.Password != nil {
		hashedPassword, err := security.HashedPassword(*dto.Password)
		if err != nil {
			return nil, fmt.Errorf("error hashing password: %w", err)
		}
		user.HashedPassword = hashedPassword
	}

	err = s.repo.Update(user)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return user, nil
}

func (s *ToDoService) Login(email, password string) (string, error) {
	user, err := s.GetByEmail(email)
	if err != nil {
		return "", nil
	}

	if err := security.CheckPassword(password, user.HashedPassword); err != nil {
		return "", p_error.ErrInvalidAccount
	}
	token, err := security.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
