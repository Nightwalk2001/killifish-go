package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo/options"
	"killifish/docs"
	"killifish/mongo"
)

type Queries struct {
	Trigger  string `query:"trigger"`
	Result   string `query:"result"`
	Page     int    `query:"page"`
	Pagesize int64  `query:"pagesize"`
}

func GetRecordings(c *fiber.Ctx) error {
	ctx := c.Context()

	id := c.Params("tank")

	q := Queries{}
	_ = c.QueryParser(&q)

	t := time.Now().Add(time.Hour * 24 * -30).Unix()

	f := M{"tank": id, "time": M{"$gte": t}}

	count, _ := mongo.Recordings.CountDocuments(ctx, f)

	ops := options.Find().SetSkip(10).SetLimit(q.Pagesize)
	cursors, _ := mongo.Recordings.Find(ctx, f, ops)
	r := make([]docs.Recording, 0)
	_ = cursors.All(ctx, &r)

	m := fiber.Map{
		"count":      count,
		"recordings": r,
	}

	return c.JSON(m)
}
