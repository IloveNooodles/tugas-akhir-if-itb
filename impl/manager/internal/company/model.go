package company

import (
	"time"

	"github.com/google/uuid"
)

type Company struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	ClusterName string    `json:"cluster_name" db:"cluster_name"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CompanyUser struct {
	Company
	UserName  string `json:"username" db:"username"`
	UserEmail string `json:"email" db:"email"`
}
