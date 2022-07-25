package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"killifish/apis"
	"killifish/config"
	"killifish/iot"
	"killifish/mongodb"
	"killifish/schedules"
)

func main() {
	conf := config.Load()
	iot.Setup(&conf)
	mongodb.Setup(&conf)
	schedules.Setup()

	defer func() {
		iot.Disconnect()
		mongodb.Disconnect()
		schedules.CleanUp()
	}()

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ReduceMemoryUsage:     true,
	})
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	apis.Setup(app)
	log.Fatal(app.Listen(":3000"))
}
