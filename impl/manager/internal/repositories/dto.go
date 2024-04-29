package repositories

type CreateRequest struct {
	Name        string `json:"name" validate:"required,printascii"`
	Description string `json:"description" validate:"required,printascii"`
	Image       string `json:"image" validate:"required,printascii"`
}
