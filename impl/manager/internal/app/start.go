package app

import (
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

	userRepo := user.NewRepository(dep.DB, dep.Logger)
	userUsecase := user.NewUsecase(dep.Logger, userRepo)
	userHandler := user.NewHandler(dep.Logger, userUsecase)
	user.RegisterRoute(userHandler, app)

	return &StartCmd{
		Server: svr,
		Logger: dep.Logger,
	}
}

func (s *StartCmd) Start() {
	err := s.Server.Start()
	if err != nil {
		s.Logger.Fatalf("error when starting server err: %s", err)
	}
}
