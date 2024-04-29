package groupdevice

import "github.com/google/uuid"

type CreateRequest struct {
	DeviceID uuid.UUID `json:"device_id" validate:"required,printascii"`
	GroupID  uuid.UUID `json:"group_id" validate:"required,printascii"`
}
