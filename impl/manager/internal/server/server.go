package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/config"
	m "github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/middleware"
	"github.com/go-playground/validator/v10"
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

func errorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if castedErr, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedErr {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s field is required", err.Field())
			case "required_if":
				report.Message = fmt.Sprintf("%s field is required when %s", err.Field(), err.Param())
			case "email":
				report.Message = fmt.Sprintf("%s is not valid email", err.Field())
			case "gte":
				report.Message = fmt.Sprintf("%s field value must be greater than '%s'", err.Field(), err.Param())
			case "lte":
				report.Message = fmt.Sprintf("%s field value must be lower than '%s'", err.Field(), err.Param())
			case "min":
				report.Message = fmt.Sprintf("%s field value must have '%s' characters or more", err.Field(), err.Param())
			case "contains":
				report.Message = fmt.Sprintf("%s field value must include '%s' character", err.Field(), err.Param())
			case "startswith":
				report.Message = fmt.Sprintf("%s field value must startswith '%s' character", err.Field(), err.Param())
			case "oneof":
				report.Message = fmt.Sprintf("%s field value must be one of '%s'", err.Field(), err.Param())
			case "excludes":
				report.Message = fmt.Sprintf("%s field value must excludes '%s' characters", err.Field(), err.Param())
			case "dive":
				report.Message = fmt.Sprintf("%s field value must not be empty", err.Field())
			}
		}
	}

	c.Logger().Error(report)
	c.JSON(http.StatusBadRequest, report)
}

func New(l *logrus.Logger, cfg config.Config) Server {
	e := echo.New()
	e.Logger.SetOutput(l.Writer())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{cfg.FEURL},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS", "HEAD"},
	}))

	e.Use(middleware.Recover())
	e.Use(m.ValidateAPIKey)

	e.HTTPErrorHandler = errorHandler

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
