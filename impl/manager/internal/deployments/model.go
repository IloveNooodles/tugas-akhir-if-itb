package deployments

import (
	"time"

	"github.com/google/uuid"
)

type Deployment struct {
	ID           uuid.UUID `json:"id" db:"id"`
	RepositoryID uuid.UUID `json:"repository_id" db:"repository_id"`
	Name         string    `json:"name" db:"name"`
	Version      string    `json:"version" db:"version"`
	Target       string    `json:"target" db:"target"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type V0Deployment struct {
	Version  string   `json:"version" validate:"required,printascii"`
	Name     string   `json:"name" validate:"required,printascii"`
	Category string   `json:"category" validate:"required,printascii"`
	Steps    []string `json:"steps"`
	RecipeID string   `json:"recipeID" validate:"omitempty,printascii"`
}

type Recipes struct {
}

type Action struct {
	Name string `json:"name" validate:"required,printascii"`
	Type string `json:"type" validate:"required,printascii"`
}

type Input struct{}

type Output struct{}
