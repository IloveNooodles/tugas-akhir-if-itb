package device

import (
	"github.com/google/uuid"
)

type CreateRequest struct {
	Name       string    `json:"name" validate:"required,printascii"`
	Type       string    `json:"type" validate:"required,printascii"`
	CompanyID  uuid.UUID `json:"company_id" validate:"required,printascii"`
	Attributes string    `json:"attributes" validate:"required,printascii"`
}
