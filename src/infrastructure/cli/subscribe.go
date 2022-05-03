package cli

import (
	"context"
	"farm/configs"
	"farm/src/FarmContext/logs"
	"farm/src/infrastructure"
	"farm/src/infrastructure/consumer"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "start listening from Mqtt Broker",
	Long:  `start listening from Mqtt Broker`,
	Run: func(cmd *cobra.Command, args []string) {
		subscribe()
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}

func subscribe() {
	client, err := consumer.NewMqttClient(configs.Config.Mqtt)
	if err != nil {
		log.Panic(err)
	}
	ctx := context.Background()
	handler := logs.NewUnmarshaller(logs.NewDeviceLogRepository(infrastructure.PostgresDBProvider))
	client.Subscribe("fleet/v1/#", 0, func(client mqtt.Client, message mqtt.Message) {
		handler.Unmarshal(ctx, logs.DeviceLogMessage{
			PayLoad: message.Payload(),
			Topic:   message.Topic(),
		})
		/*if err != nil {
			log.Errorln(err)
		}*/
	})
	<-ctx.Done()
}
