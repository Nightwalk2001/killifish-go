package main

import (
	"log"

	"killifish/api"
	"killifish/config"
	"killifish/redis"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"killifish/mongo"
	"killifish/schedules"
)

func main() {
	conf := config.Load()
	//mqttx.Setup(&conf)
	mongo.Setup(&conf)
	redis.Setup(&conf)
	schedules.Setup()

	defer func() {
		//mqttx.Disconnect()
		mongo.Disconnect()
		redis.Disconnect()
		schedules.CleanUp()
	}()

	app := fiber.New(
		fiber.Config{
			DisableStartupMessage: true,
			ReduceMemoryUsage:     true,
		},
	)

	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	api.Setup(app)
	log.Fatal(app.Listen(":3000"))
}

// 自动饲喂： 设置什么时候投喂，数量
// 立刻饲喂
// 按钮
