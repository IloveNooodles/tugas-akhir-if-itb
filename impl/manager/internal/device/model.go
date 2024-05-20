package device

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Type      string    `json:"type" db:"type"`
	CompanyID uuid.UUID `json:"company_id" db:"company_id"`
	Labels    string    `json:"labels" db:"labels"`
	NodeName  string    `json:"node_name" db:"node_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type GroupDetail struct {
	ID        uuid.UUID `json:"id" db:"id"`
	GroupID   uuid.UUID `json:"group_id" db:"group_id"`
	GroupName string    `json:"name" db:"name"`
}
