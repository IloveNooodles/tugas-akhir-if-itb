package groups

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

func (u *Usecase) Create(ctx context.Context, d Group) (Group, error) {
	return u.Repo.Create(ctx, d)
}

func (u *Usecase) GetByID(ctx context.Context, id uuid.UUID) (Group, error) {
	return u.Repo.GetByID(ctx, id)
}