package history

import (
	"time"

	"github.com/google/uuid"
)

type Histories struct {
	ID           uuid.UUID `json:"id" db:"id"`
	DeviceID     uuid.UUID `json:"device_id" db:"device_id"`
	ImageID      uuid.UUID `json:"image_id" db:"image_id"`
	DeploymentID uuid.UUID `json:"deployment_id" db:"deployment_id"`
	CompanyID    uuid.UUID `json:"company_id" db:"company_id"`
	Status       string    `json:"status" db:"status"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
