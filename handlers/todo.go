package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"killifish/docs"
	"killifish/mongo"
)

func GetTodos(c *fiber.Ctx) error {
	ctx := c.Context()
	name := c.GetRespHeader("User-Name")
	f := M{"creator": name}

	cursor, e1 := mongo.Todos.Find(ctx, f)
	if e1 != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("查找错误")
	}
	todos := make([]docs.Todo, 0)
	if e2 := cursor.All(ctx, &todos); e2 != nil {
		return c.Status(fiber.StatusResetContent).JSON("序列化错误")
	}

	return c.JSON(todos)
}

func InsertTodo(c *fiber.Ctx) error {
	name := c.GetRespHeader("User-Name")
	t := docs.Todo{
		Id:       primitive.NewObjectID().Hex(),
		CreateAt: time.Now().Unix(),
		Creator:  name,
	}
	_ = c.BodyParser(&t)

	r, _ := mongo.Todos.InsertOne(c.Context(), t)
	return c.JSON(r)
}
