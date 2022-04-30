package cli

import (
	"farm/src/FarmContext/farm"
	"farm/src/FarmContext/users"
	"farm/src/infrastructure"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var createObject = &cobra.Command{
	Use:   "create [src]",
	Short: "Run farm for create db objects.",
	Long:  `Run farm for create db objects`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "farm":
			dbProvider := infrastructure.PostgresDBProvider

			farmRepo := farm.NewFarmRepository(dbProvider)
			farmCHM := farm.NewCommandHandler(farmRepo)

			command := farm.CreateFarmCommand{
				FarmName: args[1],
				Production: args[2],
			}

			_, err := farmCHM.CreateFarm(cmd.Context(), command)

			if err != nil {
				log.Panic(err)
			}

			break
		case "user":
			dbProvider := infrastructure.PostgresDBProvider

			usersRepo := users.NewUserRepository(dbProvider)
			farmRepo := farm.NewFarmRepository(dbProvider)

			userCHM := users.NewCommandHandler(farmRepo, usersRepo)

			farmID,  err := strconv.ParseUint(args[3], 10, 32)
			if err != nil {
				log.Panic(err)
			}

			command := users.CreateUserCommand{
				Username: args[1],
				Password: args[2],
				FarmID: uint(farmID),
				FirstName: args[4],
				LastName: args[5],
				Phone: args[6],
				Email: args[7],

			}

			_, err = userCHM.CreateUser(cmd.Context(), command)
			if err != nil {
				log.Panic(err)
			}



		default:
			log.Panic("ops !")
		}
	},
}

func init() {
	rootCmd.AddCommand(createObject)
}
