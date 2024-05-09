package groups

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CompanyID uuid.UUID `json:"company_id" db:"company_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type DeviceDetail struct {
	ID         uuid.UUID `json:"id" db:"id"`
	DeviceID   uuid.UUID `json:"device_id" db:"device_id"`
	Name       string    `json:"name" db:"name"`
	Type       string    `json:"type" db:"type"`
	Attributes string    `json:"attributes" db:"attributes"`
}
