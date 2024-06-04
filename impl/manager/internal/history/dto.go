package history

import "github.com/google/uuid"

type CreateRequest struct {
	DeviceID     uuid.UUID `json:"device_id" validate:"required,printascii"`
	RepositoryID uuid.UUID `json:"repository_id" validate:"required,printascii"`
	DeploymentID uuid.UUID `json:"deployment_id" validate:"required,printascii"`
	Status       string    `json:"status" validate:"required,printascii"`
}

type GetAllParams struct {
	CompanyID uuid.UUID  `json:"company_id"`
	DeviceID  uuid.UUIDs `json:"device_id"`
}
