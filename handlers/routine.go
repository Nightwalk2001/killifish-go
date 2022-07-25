package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"killifish/mongodb"
)

func GetRoutine(ctx *fiber.Ctx) error {
	context := ctx.Context()
	routines := mongodb.Routines
	r := Routine{}
	filter := bson.M{}
	if err := routines.FindOne(context, filter).Decode(&r); err != nil {
		fmt.Println(err)
		if err == mongo.ErrNoDocuments {
			newRoutine := NewRoutine()
			_, _ = routines.InsertOne(context, newRoutine)
			return ctx.JSON(newRoutine)
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON("服务端错误")
	}
	return ctx.JSON(r)
}

func UpdateRoutine(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	m := fiber.Map{}
	_ = json.Unmarshal(ctx.Body(), &m)
	update := bson.M{"$set": m}
	res, _ := mongodb.Routines.UpdateByID(ctx.Context(), id, update)

	return ctx.JSON(res)
}
