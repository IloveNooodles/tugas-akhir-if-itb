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
	q := `INSERT INTO devices (name, type, company_id, attributes) VALUES ($1, $2, $3, $4) RETURNING *`
	err := r.DB.GetContext(ctx, &device, q, d.Name, d.Type, d.CompanyID, d.Attributes)

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

func (r *Repository) GetByGroupID(ctx context.Context, companyID, groupID uuid.UUID) ([]Device, error) {
	devices := make([]Device, 0)
	q := `SELECT * FROM devices d JOIN 
  groupdevices gd 
    ON d.id = gd.device_id 
  WHERE gd.group_id = $1 AND gd.company_id = $2`
	err := r.DB.SelectContext(ctx, &devices, q, groupID, companyID)

	if err != nil {
		r.Logger.Errorf("error when get all devices err: %s", err)
		return devices, err
	}
	return devices, nil
}
