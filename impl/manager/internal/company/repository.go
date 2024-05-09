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

func (r *Repository) GetCompanyAndLoggedInUser(ctx context.Context, companyID, userID uuid.UUID) (CompanyUser, error) {
	companyUser := CompanyUser{}
	q := `SELECT 
  c.id, 
  c.name, 
  c.cluster_name, 
  c.created_at, 
  c.updated_at, 
  u."name" username, 
  u.email email 
FROM companies c 
JOIN users u ON c.id = u.company_id 
WHERE 
  c.id = $1 
AND 
  u.id = $2`
	err := r.DB.GetContext(ctx, &companyUser, q, companyID, userID)

	if err != nil {
		r.Logger.Errorf("error when getting companies with user id: %s, userID: %s, err: %s", companyID, userID, err)
		return companyUser, err
	}

	return companyUser, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	q := `DELETE FROM companies WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, q, id)
	return err
}
