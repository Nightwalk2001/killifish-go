package redis

import (
	goredis "github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
	"killifish/config"
)

var Redis *goredis.Client
var ReJson *rejson.Handler

func Setup(conf *config.Config) {
	opts := goredis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPass,
		DB:       0,
	}

	Redis = goredis.NewClient(&opts)
	ReJson = rejson.NewReJSONHandler()
	ReJson.SetGoRedisClient(Redis)
}

func Disconnect() {
	_ = Redis.Close()
}
