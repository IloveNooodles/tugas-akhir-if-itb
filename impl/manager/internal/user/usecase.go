package user

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

func (u *Usecase) Create(ctx context.Context, user User) (User, error) {
	return u.Repo.Create(ctx, user)
}
func (u *Usecase) GetByID(ctx context.Context, id uuid.UUID) (User, error) {
	return u.Repo.GetByID(ctx, id)
}
func (u *Usecase) GetByEmail(ctx context.Context, email string) (User, error) {
	return u.Repo.GetByEmail(ctx, email)
}
