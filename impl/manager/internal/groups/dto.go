package groups

import "github.com/google/uuid"

type CreateRequest struct {
	Name      string    `json:"name" db:"name"`
	CompanyID uuid.UUID `json:"company_id" db:"company_id"`
}
