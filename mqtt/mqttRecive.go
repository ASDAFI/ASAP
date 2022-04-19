package mqtt

import (
	"crypto/tls"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"os"
)

type IMessageReceiver interface {
	Receive(MQTT.Message) error
}

func Receive(message MQTT.Message) error {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	return nil
}

func NewMqttReceiver(topicPrefix string, receiver IMessageReceiver, mqtt MQTT.Client) *MqttReceiver {
	return &MqttReceiver{
		Mqtt:            mqtt,
		TopicPrefix:     topicPrefix,
		MessageReceiver: receiver,
		Qos:             byte(0),
	}
}

type MqttReceiver struct {
	Mqtt            MQTT.Client
	MessageReceiver IMessageReceiver
	TopicPrefix     string
	Qos             byte
}

func (r *MqttReceiver) Start(interrupt <-chan os.Signal) {

	topic := r.TopicPrefix + "#"
	token := r.Mqtt.Subscribe(topic, r.Qos, func(client MQTT.Client, message MQTT.Message) {
		//err := r.MessageReceiver.Receive(Message{
		//	Payload: message.Payload(),
		//	Topic:   message.Topic(),
		//})
		err := r.MessageReceiver.Receive(message)
		if err != nil {
			log.WithFields(log.Fields{
				"topic":   message.Topic(),
				"message": string(message.Payload()),
			}).Warning(err)
		}
	})

	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.WithField("topic", r.TopicPrefix).Info("Subscribed on MQTT TopicPrefix.")

	<-interrupt
	r.Mqtt.Unsubscribe(topic).Wait()
}

func NewMqtt(server string, clientId string, username string, password string) MQTT.Client {

	connOpts := &MQTT.ClientOptions{
		ClientID: clientId,
	}

	connOpts.AddBroker(server)

	if username != "" {
		connOpts.SetUsername(username)
	}

	if password != "" {
		connOpts.SetPassword(password)
	}

	connOpts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		log.Panic(err)
	})

	connOpts.SetAutoReconnect(false)

	connOpts.SetTLSConfig(&tls.Config{
		InsecureSkipVerify: true,
		ClientAuth:         tls.NoClientCert,
	})

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		log.WithFields(log.Fields{
			"server":   server,
			"clientId": clientId,
		}).Info("Connected to MQTT.")

		return client
	}
	return nil
}
