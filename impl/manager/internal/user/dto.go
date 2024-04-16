package user

import "github.com/google/uuid"

type CreateRequest struct {
	Name      string    `json:"name" validate:"required,printascii"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,printascii"`
	CompanyID uuid.UUID `json:"company_id" validate:"required,printascii"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,printascii"`
}
