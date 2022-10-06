package mongo

import (
	"context"
	"time"

	"killifish/config"

	gomongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *gomongo.Client

	Persons    *gomongo.Collection
	Tanks      *gomongo.Collection
	Recordings *gomongo.Collection
	Todos      *gomongo.Collection
	Routines   *gomongo.Collection
	Operations *gomongo.Collection
	Alerts     *gomongo.Collection
)

func Setup(conf *config.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, _ = gomongo.Connect(ctx, options.Client().ApplyURI(conf.Uri))
	db := client.Database("killifish")

	Persons = db.Collection("persons")
	Tanks = db.Collection("tanks")
	Recordings = db.Collection("recordings")
	Todos = db.Collection("todos")
	Routines = db.Collection("routines")
	Operations = db.Collection("operations")
	Alerts = db.Collection("alerts")
}

func Disconnect() {
	_ = client.Disconnect(context.Background())
}
