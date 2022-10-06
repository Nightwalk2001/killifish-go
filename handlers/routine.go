package handlers

import (
	"encoding/json"
	"time"

	"killifish/docs"

	"github.com/gofiber/fiber/v2"
	gomongo "go.mongodb.org/mongo-driver/mongo"
	"killifish/mongo"
)

func GetRoutine(c *fiber.Ctx) error {

	ctx := c.Context()
	routines := mongo.Routines
	r := docs.Routine{}
	date := time.Now().Format("2006-01-02")
	f := M{"_id": M{"$gte": date}}
	if err := routines.FindOne(ctx, f).Decode(&r); err != nil {

		if err == gomongo.ErrNoDocuments {
			newRoutine := docs.NewRoutine()
			_, _ = routines.InsertOne(ctx, newRoutine)
			return c.JSON(newRoutine)
		}
		return c.Status(fiber.StatusInternalServerError).JSON("服务端错误")
	}
	return c.JSON(r)
}

func GetPastRoutine(c *fiber.Ctx) error {
	r := docs.Routine{}
	date := time.Now().Add(time.Hour * -24).Format("2006-01-02")
	f := M{"_id": date}
	if err := mongo.Routines.FindOne(c.Context(), f).Decode(&r); err != nil {

		if err == gomongo.ErrNoDocuments {
			return c.JSON("no")
		}
		return c.Status(fiber.StatusInternalServerError).JSON("服务端错误")
	}
	return c.JSON(r)
}

func UpdateRoutine(c *fiber.Ctx) error {
	id := time.Now().Format("2006-01-02")
	m := fiber.Map{}
	_ = json.Unmarshal(c.Body(), &m)

	update := M{"$set": m}
	res, _ := mongo.Routines.UpdateByID(c.Context(), id, update)

	return c.JSON(res)
}
