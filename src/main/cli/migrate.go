package cli

import (
	"farm/src/models"
	"github.com/spf13/cobra"
)

var migrate = &cobra.Command{
	Use: "migrate",
	Run: func(cobraCmd *cobra.Command, args []string) {
		models.PostgresDBProvider.MigrateDB()
	},
}

func init() {
	rootCmd.AddCommand(migrate)
}
