package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Uri      string
	Broker   string
	User     string
	Password string
	ClientId string
}

func Load() Config {
	_ = godotenv.Load()
	uri := os.Getenv("URI")
	broker := os.Getenv("BROKER")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	clientId := os.Getenv("CLIENT_ID")

	return Config{
		Uri:      uri,
		Broker:   broker,
		User:     user,
		Password: password,
		ClientId: clientId,
	}
}

const Mail string = "wangwei@nibs.ac.cn"
