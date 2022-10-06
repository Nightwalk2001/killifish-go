package mqttx

import (
	"context"
	"encoding/json"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"killifish/mongo"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	payload := msg.Payload()
	recording := Recording{}
	_ = json.Unmarshal(payload, &recording)

	_, _ = mongo.Recordings.InsertOne(ctx, recording)

}

type Payload struct {
	Tank    string `json:"tank"`
	Message string `json:"message"`
}

var e mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	//m := Payload{}
	//_ = json.Unmarshal(msg.Payload(), &m)
	//alert := docs.Alert{
	//	Id:      primitive.ObjectID{}.Hex(),
	//	Tank:    m.Tank,
	//	Message: m.Message,
	//	Time:    time.Now().Format(time.RFC3339),
	//}
}
