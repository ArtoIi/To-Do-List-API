package todoDTO

type DTO struct {
	Title       string `json:"title,omitempty" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
}

type Filter struct {
	Page  int
	Limit int
}
