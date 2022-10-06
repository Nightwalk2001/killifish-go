package mqttx

import (
	"killifish/config"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Recording struct {
	Name string `json:"name"`

	Id      int  `json:"id"`
	Count   int  `json:"count"`
	Succeed bool `json:"succeed"`
}

var Client mqtt.Client

func Setup(conf *config.Config) {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().
		AddBroker(conf.Broker).
		SetUsername(conf.User).
		SetPassword(conf.Password).
		SetClientID(conf.ClientId)

	opts.SetKeepAlive(1 * time.Hour)
	opts.SetPingTimeout(10 * time.Minute)

	opts.SetDefaultPublishHandler(f)

	Client = mqtt.NewClient(opts)
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	Client.Subscribe("feeding-result", 1, f)
	Client.Subscribe("errors", 1, e)
}

func Disconnect() {
	Client.Disconnect(1)
}
