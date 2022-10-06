package api

import (
	"github.com/gofiber/fiber/v2"
	"killifish/handlers"
)

func Setup(app *fiber.App) {
	app.Use(auth)

	api := app.Group("/api")

	api.Get("/names", handlers.GetPersonNames)
	api.Get("/persons", handlers.GetPersons)

	api.Post("/signin", handlers.Signin)
	api.Post("/signup", handlers.Signup)

	api.Get("/tanks/stat", handlers.StatTanks)
	api.Get("/tanks/stats", handlers.StatAllTanks)

	api.Get("/tank/:id", handlers.GetTank)
	api.Post("/tanks", handlers.GetTanks)
	api.Post("/tanks/all", handlers.GetAllTanks)
	api.Post("/tank", handlers.InsertTank)
	api.Post("/tanks", handlers.InsertTanks)
	api.Put("/tank/:id", handlers.UpdateTank)
	api.Put("/tanks", handlers.UpdateTanks)
	api.Delete("/tank/:id", handlers.DeleteTank)
	api.Delete("/tanks", handlers.DeleteTanks)

	api.Get("/recordings/:id", handlers.GetRecordings)

	api.Get("/todos", handlers.GetTodos)
	api.Post("/todo", handlers.InsertTodo)

	api.Get("/routine", handlers.GetRoutine)
	api.Get("/routine/past", handlers.GetPastRoutine)
	api.Put("/routine", handlers.UpdateRoutine)

	api.Get("/state", handlers.GetState)
	api.Get("/operations", handlers.GetOperations)
	api.Post("/operation", handlers.InsertOperation)
}
