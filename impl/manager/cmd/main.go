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
		},
	})

	cli.AddCommand(&cobra.Command{
		Use:   "migrate",
		Short: "Migrate database up",
		Long:  `Starting server`,
		Run: func(c *cobra.Command, _ []string) {
			m := app.NewMigrateCmd(dep)
			err := m.Up()
			if err != nil {
				l.Errorf("error when migrating err: %s", err)
			}
		},
	})

	return cli
}
