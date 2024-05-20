package company

import (
	"context"
	"errors"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/controller"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	ErrClusterNotAvailable = errors.New("company: cluster name is not available")
)

type Usecase struct {
	Logger         *logrus.Logger
	Repo           *Repository
	KubeController *controller.KubernetesController
}

func NewUsecase(l *logrus.Logger, r *Repository, kc *controller.KubernetesController) Usecase {
	return Usecase{
		Logger:         l,
		Repo:           r,
		KubeController: kc,
	}
}

func (u *Usecase) Create(ctx context.Context, c Company) (Company, error) {
	isCtxAvailable := u.KubeController.CheckAvailableContext(c.ClusterName)
	if !isCtxAvailable {
		u.Logger.Errorf("kube: cluster is not available from kube config %s", ErrClusterNotAvailable)
		return Company{}, ErrClusterNotAvailable
	}

	return u.Repo.Create(ctx, c)
}

func (u *Usecase) GetByID(ctx context.Context, id uuid.UUID) (Company, error) {
	return u.Repo.GetByID(ctx, id)
}

func (u *Usecase) GetAll(ctx context.Context) ([]Company, error) {
	return u.Repo.GetAll(ctx)
}

func (u *Usecase) GetCompanyAndLoggedInUser(ctx context.Context, companyID, userID uuid.UUID) (CompanyUser, error) {
	return u.Repo.GetCompanyAndLoggedInUser(ctx, companyID, userID)
}

func (u *Usecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.Repo.Delete(ctx, id)
}
