package cli

import (
	"farm/src/models"
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
			farmID, err := models.CreateFarm(args[1])
			if err != nil {
				log.Println(err)
				break
			}
			err = models.CreateUser(args[2], args[3], farmID)
			if err != nil {
				log.Println(err)
				break
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
