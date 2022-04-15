package cli

import (
	"farm/src/infrastructure/grpc_server"
	"github.com/spf13/cobra"
	"log"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving farm service.",
	Long:  `Serving farm service.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	go grpc_server.RunInsecureGRPCServer()
	err := grpc_server.RunInsecureHTTPServer()
	if err != nil {
		log.Fatal(err)
	}
}
