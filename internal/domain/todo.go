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
	Post(user *ToDo) (*ToDo, error)
	GetUserId(id int) ([]*ToDo, error)
	GetId(id int) (*ToDo, error)
	Update(user *ToDo) (*ToDo, error)
	Delete(id int) error
}
