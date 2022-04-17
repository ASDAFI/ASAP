package cli

import (
	"farm/src/FarmContext/farm"
	"farm/src/FarmContext/users"
	"farm/src/infrastructure"
	"github.com/spf13/cobra"
	"log"
)

var createObject = &cobra.Command{
	Use:   "create [src]",
	Short: "Run farm for create db objects.",
	Long:  `Run farm for create db objects`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "user":
			farmCHM := farm.NewCommandHandler(nil)
			farm, err := farmCHM.Create(cmd.Context(), args[1])
			if err != nil {
				log.Panic(err)
			}
			userCMH := users.NewCommandHandler(nil, users.NewUserRepository(infrastructure.PostgresDBProvider))
			if _, err := userCMH.Create(cmd.Context(), users.CreateUserCommand{
				Username: args[2],
				Password: args[3],
				FarmID:   farm.ID,
			}); err != nil {
				log.Panic(err)
			}
			break
		default:
			log.Panic("ops !")
		}
	},
}

func init() {
	rootCmd.AddCommand(createObject)
}
