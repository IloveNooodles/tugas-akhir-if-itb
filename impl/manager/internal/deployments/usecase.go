package deployments

import (
	"context"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/controller"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Usecase struct {
	Logger *logrus.Logger
	Repo   *Repository
	kc     *controller.KubernetesController
}

func NewUsecase(l *logrus.Logger, r *Repository, kc *controller.KubernetesController) Usecase {
	return Usecase{
		Logger: l,
		Repo:   r,
		kc:     kc,
	}
}

func (u *Usecase) Create(ctx context.Context, d Deployment) (Deployment, error) {
	return u.Repo.Create(ctx, d)
}

func (u *Usecase) GetByID(ctx context.Context, id uuid.UUID) (Deployment, error) {
	return u.Repo.GetByID(ctx, id)
}

func (u *Usecase) GetAll(ctx context.Context) ([]Deployment, error) {
	return u.Repo.GetAll(ctx)
}

func (u *Usecase) Deploy(ctx context.Context) {

}
