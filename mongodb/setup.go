package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"killifish/config"
)

var (
	client     *mongo.Client
	Recordings *mongo.Collection
	Routines   *mongo.Collection
)

func Setup(conf *config.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(conf.Uri))
	db := client.Database("killifish")

	Recordings = db.Collection("recordings")
	Routines = db.Collection("routines")
}

func Disconnect() {
	_ = client.Disconnect(context.TODO())
}
