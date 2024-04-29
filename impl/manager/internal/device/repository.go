package device

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

func (r *Repository) Create(ctx context.Context, d Device) (Device, error) {
	device := Device{}
	q := `INSERT INTO devices (name, type, company_id, attributes, node_name) VALUES ($1, $2, $3, $4, $5) RETURNING *`
	err := r.DB.GetContext(ctx, &device, q, d.Name, d.Type, d.CompanyID, d.Attributes, d.NodeName)

	if err != nil {
		r.Logger.Errorf("error when creating devices %v, err: %s", device, err)
		return device, err
	}

	return device, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (Device, error) {
	device := Device{}
	q := `SELECT * FROM devices WHERE id = $1`
	err := r.DB.GetContext(ctx, &device, q, id)

	if err != nil {
		r.Logger.Errorf("error when get device with id: %s, err: %s", id, err)
		return device, err
	}

	return device, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Device, error) {
	devices := make([]Device, 0)
	q := `SELECT * FROM devices`
	err := r.DB.SelectContext(ctx, &devices, q)

	if err != nil {
		r.Logger.Errorf("error when get all devices err: %s", err)
		return devices, err
	}
	return devices, nil
}

func (r *Repository) GetGroups(ctx context.Context, companyID, deviceID uuid.UUID) ([]GroupDetail, error) {
	devices := make([]GroupDetail, 0)
	q := `select g.id group_id, g."name" group_name
  from groupdevices gd 
  join "groups" g on gd.group_id = g.id 
  where gd.device_id = $1 AND gd.company_id = $2`
	err := r.DB.SelectContext(ctx, &devices, q, deviceID, companyID)

	if err != nil {
		r.Logger.Errorf("error when get all devices err: %s", err)
		return devices, err
	}

	return devices, nil
}
