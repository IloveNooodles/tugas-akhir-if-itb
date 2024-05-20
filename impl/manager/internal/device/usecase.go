package device

import (
	"context"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/controller"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/util"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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

func (u *Usecase) Create(ctx context.Context, d Device) (Device, error) {
	k, v, err := util.SplitByEqual(d.Labels)
	if err != nil {
		return Device{}, err
	}

	err = u.KubeController.LabelNodes(ctx, d.NodeName, k, v)
	if err != nil {
		return Device{}, err
	}

	return Device{}, nil
	// return u.Repo.Create(ctx, d)
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
