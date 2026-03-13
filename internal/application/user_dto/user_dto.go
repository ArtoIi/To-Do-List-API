package userDTO

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateUserDTO struct {
	Id       string  `json:"id"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type LoginUserDTO struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
