package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	Logger *logrus.Logger
	DB     *sqlx.DB
}

func NewRepository(db *sqlx.DB, logger *logrus.Logger) *Repository {
	return &Repository{
		DB:     db,
		Logger: logger,
	}
}

func (r *Repository) Create(ctx context.Context, d Repositories) (Repositories, error) {
	Repositories := Repositories{}
	q := `INSERT INTO deployment_repositories (name, description, image) VALUES ($1, $2, $3) RETURNING *`
	err := r.DB.GetContext(ctx, &Repositories, q, d.Name, d.Description, d.Image)

	if err != nil {
		r.Logger.Errorf("error when creating Repositories %v, err: %s", Repositories, err)
		return Repositories, err
	}

	return Repositories, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (Repositories, error) {
	Repositories := Repositories{}
	q := `SELECT * FROM deployment_repositories WHERE id = $1`
	err := r.DB.GetContext(ctx, &Repositories, q, id)

	if err != nil {
		r.Logger.Errorf("error when get repositories with id: %s, err: %s", id, err)
		return Repositories, err
	}

	return Repositories, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Repositories, error) {
	listRepositories := make([]Repositories, 0)
	q := `SELECT * FROM deployment_repositories`
	err := r.DB.SelectContext(ctx, &listRepositories, q)

	if err != nil {
		r.Logger.Errorf("error when getting list of groups err: %s", err)
		return listRepositories, err
	}

	return listRepositories, nil
}
