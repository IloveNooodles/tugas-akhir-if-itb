package deployments

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

func (r *Repository) Create(ctx context.Context, d Deployment) (Deployment, error) {
	Deployment := Deployment{}
	q := `INSERT INTO deployments (repository_id, name, version, target) VALUES ($1, $2, $3, $4) RETURNING *`
	err := r.DB.GetContext(ctx, &Deployment, q, d.RepositoryID, d.Name, d.Version, d.Target)

	if err != nil {
		r.Logger.Errorf("error when creating deployment %v, err: %s", Deployment, err)
		return Deployment, err
	}

	return Deployment, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (Deployment, error) {
	Deployment := Deployment{}
	q := `SELECT * FROM deployments WHERE id = $1`
	err := r.DB.GetContext(ctx, &Deployment, q, id)

	if err != nil {
		r.Logger.Errorf("error when get deployment with id: %s, err: %s", id, err)
		return Deployment, err
	}

	return Deployment, nil
}

func (r *Repository) GetAll(ctx context.Context) ([]Deployment, error) {
	listDeployment := make([]Deployment, 0)
	q := `SELECT * FROM deployments`
	err := r.DB.SelectContext(ctx, &listDeployment, q)

	if err != nil {
		r.Logger.Errorf("error when getting list of groups err: %s", err)
		return listDeployment, err
	}

	return listDeployment, nil
}

func (r *Repository) GetAllByCompanyID(ctx context.Context, companyID uuid.UUID) ([]Deployment, error) {
	listDeployment := make([]Deployment, 0)
	q := `SELECT * FROM deployments WHERE company_id = $1`
	err := r.DB.SelectContext(ctx, &listDeployment, q, companyID)

	if err != nil {
		r.Logger.Errorf("error when getting list of groups err: %s", err)
		return listDeployment, err
	}

	return listDeployment, nil
}

func (r *Repository) GetDeploymentWithRepository(ctx context.Context, id uuid.UUID) (DeploymentWithRepository, error) {
	dr := DeploymentWithRepository{}
	q := `select 
    d.id, d."name", 
    d."version", 
    d.created_at, 
    d.updated_at, 
    d.target,  
    dr."name" repository_name,
    dr.description repository_description,
    dr.image repository_image
  from deployments d 
  join deployment_repositories dr 
  on d.repository_id = dr.id
  WHERE d.id = $1`
	err := r.DB.GetContext(ctx, &dr, q, id)

	if err != nil {
		r.Logger.Errorf("error when getting list of groups err: %s", err)
		return dr, err
	}

	return dr, nil

}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	q := `DELETE FROM deployments WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, q, id)
	return err
}
