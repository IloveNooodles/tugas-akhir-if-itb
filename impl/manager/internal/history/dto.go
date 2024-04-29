package history

import "github.com/google/uuid"

type CreateRequest struct {
	DeviceID     uuid.UUID `json:"device_id" validate:"required,printascii"`
	ImageID      uuid.UUID `json:"image_id" validate:"required,printascii"`
	DeploymentID uuid.UUID `json:"deployment_id" validate:"required,printascii"`
	Status       string    `json:"status" validate:"required,printascii"`
}
