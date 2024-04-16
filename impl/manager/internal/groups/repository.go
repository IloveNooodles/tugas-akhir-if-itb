package groups

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

func (r *Repository) Create(ctx context.Context, d Group) (Group, error) {
	group := Group{}
	q := `INSERT INTO groups (name, company_id) VALUES ($1, $2) RETURNING *`
	err := r.DB.GetContext(ctx, &group, q, d.Name, d.CompanyID)

	if err != nil {
		r.Logger.Errorf("error when creating group %v, err: %s", group, err)
		return group, err
	}

	return group, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (Group, error) {
	group := Group{}
	q := `SELECT * FROM groups WHERE id = $1`
	err := r.DB.GetContext(ctx, &group, q, id)

	if err != nil {
		r.Logger.Errorf("error when get group with id: %s, err: %s", id, err)
		return group, err
	}

	return group, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Group, error) {
	groups := make([]Group, 0)
	q := `SELECT * FROM groups`
	err := r.DB.SelectContext(ctx, &groups, q)

	if err != nil {
		r.Logger.Errorf("error when getting list of groups err: %s", err)
		return groups, err
	}

	return groups, nil
}

func (r *Repository) GetByDeviceID(ctx context.Context, companyID, deviceID uuid.UUID) ([]Group, error) {
	groups := make([]Group, 0)
	q := `SELECT * FROM groups g JOIN
  groupdevices gd
    ON g.id = gd.group_id
  WHERE gd.device_id = $1 AND gd.company_id = $2`
	err := r.DB.SelectContext(ctx, &groups, q, deviceID, companyID)

	if err != nil {
		r.Logger.Errorf("error when getting list of groups err: %s", err)
		return groups, err
	}

	return groups, nil
}
