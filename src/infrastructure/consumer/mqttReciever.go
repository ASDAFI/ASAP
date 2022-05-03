package consumer

import (
	"farm/configs"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"os/user"
	"time"
)

func NewMqttClient(mqttConfig configs.MqttConfigs) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", mqttConfig.Host, mqttConfig.Port))

	mqttUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	opts.SetClientID(fmt.Sprintf("%s:%s", mqttUser.Username, time.Now().String()))
	opts.SetUsername(mqttConfig.Username)
	opts.SetPassword(mqttConfig.Password)

	opts.OnConnect = func(client mqtt.Client) {
		log.WithFields(log.Fields{
			"username": mqttConfig.Username,
			"host":     mqttConfig.Host,
		}).Infoln("Connecting to Mqtt Broker ...")
	}

	client := mqtt.NewClient(opts)
	log.WithFields(log.Fields{
		"username": mqttConfig.Username,
		"host":     mqttConfig.Host,
	}).Infoln("Try To Connect ...")

	time.Sleep(time.Second * 5) // waiting for logrus to log

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return client, nil

}
