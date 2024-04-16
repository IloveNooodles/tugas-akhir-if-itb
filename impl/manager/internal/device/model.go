package device

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Type       string    `json:"type" db:"type"`
	CompanyID  uuid.UUID `json:"company_id" db:"company_id"`
	Attributes string    `json:"attributes" db:"attributes"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
