package domain

<<<<<<< Updated upstream
type CreateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type User struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

type ToDoRepository interface {
	Register(DTO CreateUserDTO) error
=======
import "time"

type ToDo struct {
	ID          int       `json:"id,omitempty"`
	UserID      int       `json:"user_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type ToDoRepository interface {
	Create(todo *ToDo) error
	Update(todo *ToDo) error
	Delete(todo *ToDo) error
	GetId(id string) (*ToDo, error)
	GetUser(userId string) ([]*ToDo, error)
>>>>>>> Stashed changes
}
