package company

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

func (u *Usecase) Create(ctx context.Context, c Company) (Company, error) {
	return u.Repo.Create(ctx, c)
}

func (u *Usecase) GetByID(ctx context.Context, id uuid.UUID) (Company, error) {
	return u.Repo.GetByID(ctx, id)
}

func (u *Usecase) GetAll(ctx context.Context) ([]Company, error) {
	return u.Repo.GetAll(ctx)
}
