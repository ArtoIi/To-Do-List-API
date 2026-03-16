package domain

import "time"

type ToDo struct {
	ID          int       `json:"id,omitempty"`
	UserID      int       `json:"name,omitempty"`
	Title       string    `json:"email,omitempty"`
	Description string    `json:"hashed_password,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type ToDoRepository interface {
	Post(user *ToDo) error
	GetUserId(userID int, limit, offset int) ([]*ToDo, int, error)
	GetId(id int) (*ToDo, error)
	Update(user *ToDo) (*ToDo, error)
	Delete(id int) error
}
