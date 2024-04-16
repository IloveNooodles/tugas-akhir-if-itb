package groupdevice

import (
	"time"

	"github.com/google/uuid"
)

type GroupDevice struct {
	ID        uuid.UUID `json:"id" db:"id"`
	GroupID   uuid.UUID `json:"group_id" db:"group_id"`
	DeviceID  uuid.UUID `json:"device_id" db:"device_id"`
	CompanyID uuid.UUID `json:"company_id" db:"company_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
