package repositories

type CreateRequest struct {
	Name        string `json:"name" validate:"required,printascii,min=8"`
	Description string `json:"description" validate:"required,printascii,min=8"`
	Image       string `json:"image" validate:"required,printascii"`
}
