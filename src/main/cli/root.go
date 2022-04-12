package cli

import (
	"farm/configs"
	"farm/src/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string

	rootCmd = &cobra.Command{
		Use:   "farm",
		Short: "Farm monitoring service.",
		Long:  `Farm monitoring service.`,
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "farm-configs", "config file")
	rootCmd.PersistentFlags().StringP("author", "a", "Abolfazl Taheri", "taheri.abolfazl@gmail.com")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	viper.Unmarshal(&configs.Config)
	log.Info("Configuration initialized!")

	if configs.Config.Credential.TokenSecret == "" {
		log.Fatal("There is no token secret in config file\n")
	}

	dbProvider, err := models.CreateDBProvider(configs.Config.Database)
	if err != nil {
		log.Fatalf("Fatal error on create db: %s \n", err)
	}
	models.PostgresDBProvider = dbProvider
}
