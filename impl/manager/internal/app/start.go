package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/company"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/controller"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/deployments"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/device"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/groupdevice"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/groups"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/handler"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/history"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/repositories"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/server"
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/user"
	"github.com/sirupsen/logrus"
)

type StartCmd struct {
	Server server.Server
	Logger *logrus.Logger
}

func NewStartCmd(dep *Dep) *StartCmd {
	svr := server.New(dep.Logger, dep.Config)
	app := svr.App()

	kc, err := controller.New(dep.Logger)
	if err != nil {
		panic(err)
	}

	commonRoute := handler.New()
	handler.RegisterRoute(app, commonRoute)

	companyRepo := company.NewRepository(dep.DB, dep.Logger)
	companyUsecase := company.NewUsecase(dep.Logger, companyRepo, kc)
	companyHandler := company.NewHandler(dep.Logger, companyUsecase)
	company.RegisterRoute(companyHandler, app)

	userRepo := user.NewRepository(dep.DB, dep.Logger)
	userUsecase := user.NewUsecase(dep.Logger, userRepo)
	userHandler := user.NewHandler(dep.Logger, userUsecase, companyUsecase)
	user.RegisterRoute(userHandler, app)

	deviceRepo := device.NewRepository(dep.DB, dep.Logger)
	deviceUsecase := device.NewUsecase(dep.Logger, deviceRepo, kc)
	deviceHandler := device.NewHandler(dep.Logger, deviceUsecase, companyUsecase)
	device.RegisterRoute(deviceHandler, app)

	groupRepo := groups.NewRepository(dep.DB, dep.Logger)
	groupUsecase := groups.NewUsecase(dep.Logger, groupRepo)
	groupHandler := groups.NewHandler(dep.Logger, groupUsecase, companyUsecase)
	groups.RegisterRoute(groupHandler, app)

	groupDeviceRepo := groupdevice.NewRepository(dep.DB, dep.Logger)
	groupDeviceUsecase := groupdevice.NewUsecase(dep.Logger, groupDeviceRepo)
	groupDeviceHandler := groupdevice.NewHandler(dep.Logger, groupDeviceUsecase, companyUsecase)
	groupdevice.RegisterRoute(groupDeviceHandler, app)

	repositoriesRepo := repositories.NewRepository(dep.DB, dep.Logger)
	repositoriesUsecase := repositories.NewUsecase(dep.Logger, repositoriesRepo)
	repositoriesHandler := repositories.NewHandler(dep.Logger, repositoriesUsecase, companyUsecase)
	repositories.RegisterRoute(repositoriesHandler, app)

	historyRepo := history.NewRepository(dep.DB, dep.Logger)
	historyUsecase := history.NewUsecase(dep.Logger, historyRepo)
	historyHandler := history.NewHandler(dep.Logger, historyUsecase, companyUsecase)
	history.RegisterRoute(historyHandler, app)

	deploymentRepo := deployments.NewRepository(dep.DB, dep.Logger)
	deploymentUsecase := deployments.NewUsecase(dep.Logger, deploymentRepo, kc)
	deploymentHandler := deployments.NewHandler(dep.Logger, deploymentUsecase, companyUsecase, historyUsecase, deviceUsecase, groupUsecase, dep.Config)
	deployments.RegisterRoute(deploymentHandler, app)

	return &StartCmd{
		Server: svr,
		Logger: dep.Logger,
	}
}

func (s *StartCmd) Start() {
	sigs := make(chan os.Signal, 1)

	go func() {
		if err := s.Server.Start(); err != nil && errors.Is(err, http.ErrServerClosed) {
			s.Logger.Errorf("error when starting server err: %s", err)
		}
	}()

	signal.Notify(sigs, os.Interrupt)
	<-sigs

	s.Logger.Info("attempting graceful shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Server.Stop(ctx); err != nil {
		s.Logger.Error(err)
	}

	s.Logger.Info("graceful shutdown completed")

}
