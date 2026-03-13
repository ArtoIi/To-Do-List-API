package todoDTO

type CreateToDoDTO struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
