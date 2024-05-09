package device

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

func (u *Usecase) Create(ctx context.Context, d Device) (Device, error) {
	return u.Repo.Create(ctx, d)
}

func (u *Usecase) GetByID(ctx context.Context, id uuid.UUID) (Device, error) {
	return u.Repo.GetByID(ctx, id)
}

func (u *Usecase) GetAll(ctx context.Context) ([]Device, error) {
	return u.Repo.GetAll(ctx)
}

func (u *Usecase) GetGroups(ctx context.Context, companyID, deviceID uuid.UUID) ([]GroupDetail, error) {
	return u.Repo.GetGroups(ctx, companyID, deviceID)
}

func (u *Usecase) GetAllByCompanyID(ctx context.Context, companyID uuid.UUID) ([]Device, error) {
	return u.Repo.GetAllByCompanyID(ctx, companyID)
}

func (u *Usecase) Delete(ctx context.Context, id uuid.UUID) error {
  return u.Repo.Delete(ctx, id)
}