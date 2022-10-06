package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Uri       string
	RedisAddr string
	RedisPass string
	Broker    string
	User      string
	Password  string
	ClientId  string
}

func Load() Config {
	_ = godotenv.Load(".env.dev")
	uri := os.Getenv("URI")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASS")
	broker := os.Getenv("BROKER")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	clientId := os.Getenv("CLIENT_ID")

	return Config{
		Uri:       uri,
		RedisAddr: redisAddr,
		RedisPass: redisPass,
		Broker:    broker,
		User:      user,
		Password:  password,
		ClientId:  clientId,
	}
}

const Mail string = "wangwei@nibs.ac.cn"

var Secret = []byte(os.Getenv("SECRET"))
