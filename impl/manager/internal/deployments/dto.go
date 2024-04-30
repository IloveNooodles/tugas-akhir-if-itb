package deployments

import "github.com/google/uuid"

type CreateRequest struct {
	Name         string    `json:"name" validate:"required,printascii"`
	Version      string    `json:"version" validate:"required,printascii"`
	Target       string    `json:"target" validate:"required,printascii"`
	RepositoryID uuid.UUID `json:"repository_id" validate:"required,printascii"`
}

type DeploymentRequest struct {
	DeploymentIDs uuid.UUIDs `json:"deployment_ids" validate:"required,dive,printascii"`
}
