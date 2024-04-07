package app

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/config"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Dep struct {
	Config config.Config
	Logger *logrus.Logger
	DB     *sqlx.DB
}

func NewDep() *Dep {
	cfg := config.New()

	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			_, after, _ := strings.Cut(frame.File, "impl/")
			fileName := fmt.Sprintf("%s:%d", after, frame.Line)

			s := strings.Split(frame.Function, ".")
			funcname := s[len(s)-1]

			return funcname, fileName
		},
	})

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDB, cfg.PostgresSSLMode)
	db, err := sqlx.Connect("postgres", dbUrl)

	if err != nil {
		l.Fatalf("error when connecting db, url: %s, err: %s", dbUrl, err)
	}

	return &Dep{
		Config: cfg,
		Logger: l,
		DB:     db,
	}
}
