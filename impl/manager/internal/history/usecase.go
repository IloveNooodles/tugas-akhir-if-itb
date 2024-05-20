package history

import (
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	Logger *logrus.Logger
	Repo   *Repository
}

func NewUsecase(l *logrus.Logger, r *Repository) Usecase {
	return Usecase{
		Logger: l,
		Repo:   r,
	}
}

func (u *Usecase) Create(ctx context.Context, d Histories) (Histories, error) {
	return u.Repo.Create(ctx, d)
}

func (u *Usecase) GetByID(ctx context.Context, id uuid.UUID) (Histories, error) {
	return u.Repo.GetByID(ctx, id)
}

func (u *Usecase) GetAll(ctx context.Context) ([]Histories, error) {
	return u.Repo.GetAll(ctx)
}

func (u *Usecase) GetAllByCompanyID(ctx context.Context, companyID uuid.UUID) ([]Histories, error) {
	return u.Repo.GetAllByCompanyID(ctx, companyID)
}

func (u *Usecase) UpdateStatusById(ctx context.Context, ID uuid.UUID, status string) (Histories, error){
  return u.Repo.UpdateStatusById(ctx, ID, status)
}
