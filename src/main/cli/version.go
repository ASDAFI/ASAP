package cli

import (
	"log"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version info",
	Long:  `Prints the current version info.`,
	Run: func(cmd *cobra.Command, args []string) {
		logVersion()
	},
}
var (
	dumpConfig bool
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func logVersion() {
	log.Printf("1.0.0")
}
