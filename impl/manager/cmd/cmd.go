package cmd

import (
	"github.com/IloveNooodles/tugas-akhir-if-itb/impl/manager/internal/app"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cli := &cobra.Command{}
	dep := app.NewDep()
	l := dep.Logger

	cli.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Starting server",
		Long:  `Starting server`,
		Run: func(c *cobra.Command, _ []string) {
			startCmd := app.NewStartCmd(dep)
			startCmd.Start()
		},
	})

	cli.AddCommand(&cobra.Command{
		Use:   "migrate-up",
		Short: "Migrate database up",
		Long:  `Starting server`,
		Run: func(c *cobra.Command, _ []string) {
			m := app.NewMigrateCmd(dep)
			err := m.Up()
			if err != nil {
				l.Errorf("error when migrating err: %s", err)
			}
			l.Info("success migrate up")
		},
	})

	cli.AddCommand(&cobra.Command{
		Use:   "migrate-down",
		Short: "Migrate database down",
		Long:  `Starting server`,
		Run: func(c *cobra.Command, _ []string) {
			m := app.NewMigrateCmd(dep)
			err := m.Down()
			if err != nil {
				l.Errorf("error when migrating err: %s", err)
			}
			l.Info("success migrate down")
		},
	})

	return cli
}
