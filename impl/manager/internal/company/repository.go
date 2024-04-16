package company

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

func (r *Repository) Create(ctx context.Context, c Company) (Company, error) {
	company := Company{}
	q := `INSERT INTO companies (name, cluster_name) VALUES ($1, $2) RETURNING *`
	err := r.DB.GetContext(ctx, &company, q, c.Name, c.ClusterName)

	if err != nil {
		r.Logger.Errorf("error when creating company %v, err: %s", company, err)
		return company, err
	}

	return company, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (Company, error) {
	company := Company{}
	q := `SELECT * FROM companies WHERE id = $1`
	err := r.DB.GetContext(ctx, &company, q, id)

	if err != nil {
		r.Logger.Errorf("error when get company with id: %s, err: %s", id, err)
		return company, err
	}

	return company, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Company, error) {
	companies := make([]Company, 0)
	q := `SELECT * FROM companies`
	err := r.DB.SelectContext(ctx, &companies, q)

	if err != nil {
		r.Logger.Errorf("error when getting companies list err: %s", err)
		return companies, err
	}

	return companies, err
}
