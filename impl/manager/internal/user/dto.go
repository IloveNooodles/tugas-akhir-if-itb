package user

type CreateRequest struct {
	Name     string `json:"name" validate:"required,printascii"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,printascii"`
}
