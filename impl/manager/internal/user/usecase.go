package user

import (
	"context"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/errx"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Logger.Errorf("error when hashing password err: %s", err)
		return User{}, err
	}
	user.Password = string(hashedPassword)
	return u.Repo.Create(ctx, user)
}

func (u *Usecase) GetByID(ctx context.Context, id uuid.UUID) (User, error) {
	return u.Repo.GetByID(ctx, id)
}

func (u *Usecase) Login(ctx context.Context, email, password string) (User, error) {
	user, err := u.Repo.GetByEmail(ctx, email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		u.Logger.Errorf("error when comparing password err: %s", err)
		return User{}, errx.ErrInvalidPassword
	}

	return user, err
}
