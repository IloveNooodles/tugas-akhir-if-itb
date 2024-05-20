package history

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

func (r *Repository) Create(ctx context.Context, d Histories) (Histories, error) {
	Histories := Histories{}
	q := `INSERT INTO deployment_histories (device_id, repository_id, deployment_id, company_id, status) VALUES ($1, $2, $3, $4, 'DEPLOYING') RETURNING *`
	err := r.DB.GetContext(ctx, &Histories, q, d.DeviceID, d.RepositoryID, d.DeploymentID, d.CompanyID)

	if err != nil {
		r.Logger.Errorf("error when creating histories %#v, err: %s", d, err)
		return Histories, err
	}

	return Histories, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (Histories, error) {
	Histories := Histories{}
	q := `SELECT * FROM deployment_histories WHERE id = $1`
	err := r.DB.GetContext(ctx, &Histories, q, id)

	if err != nil {
		r.Logger.Errorf("error when get Histories with id: %s, err: %s", id, err)
		return Histories, err
	}

	return Histories, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Histories, error) {
	listHistories := make([]Histories, 0)
	q := `SELECT * FROM deployment_histories`
	err := r.DB.SelectContext(ctx, &listHistories, q)

	if err != nil {
		r.Logger.Errorf("error when getting list of groups err: %s", err)
		return listHistories, err
	}

	return listHistories, nil
}

func (r *Repository) GetAllByCompanyID(ctx context.Context, companyID uuid.UUID) ([]Histories, error) {
	listHistories := make([]Histories, 0)
	q := `SELECT * FROM deployment_histories WHERE company_id = $1`
	err := r.DB.SelectContext(ctx, &listHistories, q, companyID)

	if err != nil {
		r.Logger.Errorf("error when getting list of groups err: %s", err)
		return listHistories, err
	}

	return listHistories, nil
}

func (r *Repository) UpdateStatusById(ctx context.Context, ID uuid.UUID, status string) (Histories, error) {
	history := Histories{}
	q := `UPDATE deployment_histories SET status = $1 WHERE id = $2 RETURNING *`
	err := r.DB.GetContext(ctx, &history, q, status, ID)
	if err != nil {
		r.Logger.Errorf("error when updating histories status err: %s", err)
		return history, err
	}

	return history, nil
}
