package apis

import (
	"github.com/gofiber/fiber/v2"
	"killifish/handlers"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	routine := api.Group("/routine")
	routine.Get("/", handlers.GetRoutine)
	routine.Put("/:id", handlers.UpdateRoutine)
}
