package cli

import (
	"github.com/spf13/cobra"
	"log"
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
	log.Printf("First version 1")
}
