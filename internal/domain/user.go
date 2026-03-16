package domain

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

type UserRepository interface {
	Register(user *User) error
	GetEmail(email string) (*User, error)
	GetId(id int) (*User, error)
	Update(user *User) error
	Delete(id int) error
}
