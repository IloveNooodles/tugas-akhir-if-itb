package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
)

type MigrateCmd struct {
	Migrate *migrate.Migrate
	DB      *sqlx.DB
}

// This functions will create a new instance of migratecmd
func NewMigrateCmd(dep *Dep) *MigrateCmd {
	path := "migrations"
	postgresDriver, err := postgres.WithInstance(dep.DB.DB, &postgres.Config{})
	l := dep.Logger
	if err != nil {
		l.Fatalf("error when getting postgres driver: %v", err)
	}

	fsys := os.DirFS(path)
	sourceDriver, err := iofs.New(fsys, ".")
	if err != nil {
		l.Fatalf("error when getting source driver: %v", err)
	}

	m, err := migrate.NewWithInstance("manager", sourceDriver, "manager", postgresDriver)
	if err != nil {
		l.Fatalf("error when creating migrate instance: %v", err)
	}

	return &MigrateCmd{
		Migrate: m,
		DB:      dep.DB,
	}
}

func (m *MigrateCmd) Up() error {
	err := m.Migrate.Up()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("error when migrate up: %w", err)
	}

	return nil
}

func (m *MigrateCmd) Down() error {
	err := m.Migrate.Down()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("error when migrate down: %w", err)
	}

	return nil
}
