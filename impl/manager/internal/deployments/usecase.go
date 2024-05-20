package deployments

import (
	"context"
	"fmt"

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

func (u *Usecase) GetAllByCompanyID(ctx context.Context, companyID uuid.UUID) ([]Deployment, error) {
	return u.Repo.GetAllByCompanyID(ctx, companyID)
}

func (u *Usecase) Deploy(ctx context.Context, deploymentIds uuid.UUIDs) ([]DeploymentWithRepository, []error) {
	var listError = make([]error, 0)
	var listRes = make([]DeploymentWithRepository, 0)

	for _, deploymentId := range deploymentIds {
		deployment, err := u.Repo.GetDeploymentWithRepository(ctx, deploymentId)
		if err != nil {
			err := fmt.Errorf("error when doing deployment with id: %s, err: %s", deploymentId, err)
			u.Logger.Error(err)
			listError = append(listError, err)
			continue
		}

		// TODO label selector
		// TODO function to convert database to label / match
		// TODO Validate string by , separated
		// TODO Add replica

		labels := convertTargetToMap(deployment.Target)

		p := controller.DeployParams{
			Replica: 1,
			Name:    deployment.Name,
			Image:   deployment.RepositoryImage,
			Labels:  labels,
			Targets: labels,
		}

		_, err = u.kc.Deploy(ctx, p)
		if err != nil {
			err := fmt.Errorf("remote deploy: error when deploying deployments with id: %s, err: %s", deploymentId, err)
			u.Logger.Error(err)
			listError = append(listError, err)
		}

		listRes = append(listRes, deployment)
	}

	return listRes, listError
}

func (u *Usecase) DeleteDeploy(ctx context.Context, deploymentIds uuid.UUIDs) []error {
	var listError = make([]error, 0)

	for _, deploymentId := range deploymentIds {
		deployment, err := u.Repo.GetDeploymentWithRepository(ctx, deploymentId)
		if err != nil {
			err := fmt.Errorf("error when doing deployment with id: %s, err: %s", deploymentId, err)
			u.Logger.Error(err)
			listError = append(listError, err)
			continue
		}

		// TODO label selector
		// TODO function to convert database to label / match
		// TODO Validate string by , separated
		// TODO Add replica

		labels := convertTargetToMap(deployment.Target)

		p := controller.DeployParams{
			Replica: 1,
			Name:    deployment.Name,
			Image:   deployment.RepositoryImage,
			Labels:  labels,
			Targets: labels,
		}

		err = u.kc.Delete(ctx, p)
		if err != nil {
			err := fmt.Errorf("error when deploying deployments with id: %s, err: %s", deploymentId, err)
			u.Logger.Error(err)
			listError = append(listError, err)
		}

	}

	return listError
}

func (u *Usecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.Repo.Delete(ctx, id)
}

func (u *Usecase) CheckDeploymentStatus(ctx context.Context, deploymentName string) bool {
	return u.kc.CheckDeploymentStatus(ctx, deploymentName)
}
