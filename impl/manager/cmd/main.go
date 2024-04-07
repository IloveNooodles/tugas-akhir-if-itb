package cmd

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cli := &cobra.Command{}

	cli.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Starting server",
		Long:  `Starting server`,
		Run: func(c *cobra.Command, _ []string) {
		},
	})

	cli.AddCommand(&cobra.Command{
		Use:   "migrate-up",
		Short: "Migrate database up",
		Long:  `Starting server`,
		Run: func(c *cobra.Command, _ []string) {
		},
	})

	return cli
}
