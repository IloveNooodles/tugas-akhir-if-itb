package server

import (
	"context"
	"fmt"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type Server interface {
	Start() error
	Stop(ctx context.Context) error
	App() *echo.Echo
}

type server struct {
	config config.Config
	logger *logrus.Logger
	echo   *echo.Echo
}

func New(l *logrus.Logger, cfg config.Config) Server {
	e := echo.New()
	e.Logger.SetOutput(l.Writer())

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	return &server{
		config: cfg,
		logger: l,
		echo:   e,
	}
}

func (s *server) Start() error {
	conn := fmt.Sprintf(":%d", s.config.Port)
	return s.echo.Start(conn)
}

func (s *server) Stop(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}

func (s *server) App() *echo.Echo {
	return s.echo
}

func (s *server) Logger() *logrus.Logger {
	return s.logger
}