package device

import (
	"context"
	"fmt"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/util"
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
	q := `INSERT INTO devices (name, type, company_id, labels, node_name) VALUES ($1, $2, $3, $4, $5) RETURNING *`
	err := r.DB.GetContext(ctx, &device, q, d.Name, d.Type, d.CompanyID, d.Labels, d.NodeName)

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
	q := `select gd.id, g.id group_id, g."name"
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

func (r *Repository) GetAllByCompanyID(ctx context.Context, companyID uuid.UUID) ([]Device, error) {
	devices := make([]Device, 0)
	q := `SELECT * FROM devices WHERE company_id = $1`
	err := r.DB.SelectContext(ctx, &devices, q, companyID)

	if err != nil {
		r.Logger.Errorf("error when get all devices err: %s", err)
		return devices, err
	}
	return devices, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	q := `DELETE FROM devices WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, q, id)
	return err
}

func (r *Repository) GetAllByLabels(ctx context.Context, companyID uuid.UUID, label string) ([]Device, error) {
	devices := make([]Device, 0)
	q := `SELECT * FROM devices WHERE company_id = $1 AND label ilike '%' || $2 || '%'`
	err := r.DB.SelectContext(ctx, &devices, q, companyID, label)

	if err != nil {
		r.Logger.Errorf("error when get all devices err: %s", err)
		return devices, err
	}

	return devices, nil
}

func (r *Repository) GetAllByIDs(ctx context.Context, companyID uuid.UUID, ids uuid.UUIDs) ([]Device, error) {
	devices := make([]Device, 0)
	var q = fmt.Sprintf(`SELECT * FROM devices WHERE company_id = $1 AND id in (%s)`, util.GenerateQuerySQL(ids.Strings(), 2))

	if len(ids) == 0 {
		q = `SELECT * FROM devices WHERE company_id = $1`
	}

	args := []any{companyID}
	for _, i := range ids {
		args = append(args, i)
	}

	err := r.DB.SelectContext(ctx, &devices, q, args...)

	if err != nil {
		r.Logger.Errorf("error when get all devices err: %s", err)
		return devices, err
	}

	return devices, nil
}

func (r *Repository) GetAllByGroupIDs(ctx context.Context, companyID uuid.UUID, groupIDs uuid.UUIDs) ([]Device, error) {
	devices := make([]Device, 0)
	q := fmt.Sprintf(`
  SELECT 
    d.* FROM groupdevices gd 
  JOIN 
    devices d ON d.id = gd.device_id 
  WHERE 
    gd.company_id = $1
  AND 
    gd.group_id in (%s)`, util.GenerateQuerySQL(groupIDs.Strings(), 2))

	if len(groupIDs) == 0 {
		q = `
    SELECT 
      d.* FROM groupdevices gd 
    JOIN 
      devices d ON d.id = gd.device_id 
    WHERE 
      gd.company_id = $1`
	}

	args := []any{companyID}
	for _, i := range groupIDs {
		args = append(args, i)
	}

	err := r.DB.SelectContext(ctx, &devices, q, args...)

	if err != nil {
		r.Logger.Errorf("error when get all devices err: %s", err)
		return devices, err
	}

	return devices, nil

}
