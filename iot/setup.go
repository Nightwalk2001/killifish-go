package iot

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"killifish/config"
	"killifish/mongodb"
)

type Recording struct {
	Name    string `json:"name"`
	Id      int    `json:"id"`
	Count   int    `json:"count"`
	Succeed bool   `json:"succeed"`
}

var Client mqtt.Client

var ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

	defer cancel()
	conn := mongodb.Recordings
	payload := msg.Payload()
	recording := new(Recording)
	_ = json.Unmarshal(payload, recording)

	_, _ = conn.InsertOne(ctx, recording)

}

func Setup(conf *config.Config) {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().
		AddBroker(conf.Broker).
		SetUsername(conf.User).
		SetPassword(conf.Password).
		SetClientID(conf.ClientId)

	opts.SetKeepAlive(60 * time.Second)

	opts.SetDefaultPublishHandler(f)

	Client = mqtt.NewClient(opts)
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	Client.Subscribe("feeding-result", 1, f)
}

func Disconnect() {
	Client.Disconnect(1)
}
