package deployments

import "github.com/google/uuid"

type CreateRequest struct {
	Name         string    `json:"name" validate:"required,printascii,min=8"`
	Version      string    `json:"version" validate:"required,printascii,startswith=v"`
	Target       string    `json:"target" validate:"required,printascii,contains=="`
	RepositoryID uuid.UUID `json:"repository_id" validate:"required,printascii"`
}

type DeploymentRequest struct {
	DeploymentIDs uuid.UUIDs `json:"deployment_ids" validate:"required,dive,printascii"`
}
