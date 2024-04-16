package groups

import "github.com/google/uuid"

type Group struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CompanyID uuid.UUID `json:"company_id" db:"company_id"`
}
