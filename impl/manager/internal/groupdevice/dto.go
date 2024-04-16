package groupdevice

import "github.com/google/uuid"

type CreateRequest struct {
	DeviceID  uuid.UUID `json:"device_id" db:"device_id"`
	GroupID   uuid.UUID `json:"group_id" db:"group_id"`
	CompanyID uuid.UUID `json:"company_id" db:"company_id"`
}
