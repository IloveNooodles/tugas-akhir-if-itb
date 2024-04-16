package groupdevice

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

func (r *Repository) Create(ctx context.Context, gu GroupDevice) (GroupDevice, error) {
	groupDevice := GroupDevice{}
	q := `INSERT INTO groupdevices (group_id, device_id, company_id) VALUES ($1, $2, $3) RETURNING *`
	err := r.DB.GetContext(ctx, &groupDevice, q, gu.GroupID, gu.DeviceID, gu.CompanyID)

	if err != nil {
		r.Logger.Errorf("error when creating group_device %v, err: %s", groupDevice, err)
		return groupDevice, err
	}

	return groupDevice, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (GroupDevice, error) {
	groupDevice := GroupDevice{}
	q := `SELECT * FROM groupdevices WHERE id = $1`
	err := r.DB.GetContext(ctx, &groupDevice, q, id)

	if err != nil {
		r.Logger.Errorf("error when get group with id: %s, err: %s", id, err)
		return groupDevice, err
	}

	return groupDevice, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]GroupDevice, error) {
	groupDevices := make([]GroupDevice, 0)
	q := `SELECT * FROM groupdevices`
	err := r.DB.SelectContext(ctx, &groupDevices, q)

	if err != nil {
		r.Logger.Errorf("error when getting all group devices err: %s", err)
		return groupDevices, err
	}

	return groupDevices, nil
}
