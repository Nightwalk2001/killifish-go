package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"killifish/docs"
	"killifish/mongo"
	"killifish/redis"
)

const (
	IncrBorn      = "ib"
	IncrKilled    = "ik"
	MutateCurrent = "mc"
	MutateTotal   = "mt"
)

func GetState(c *fiber.Ctx) error {
	//s := make(map[string]interface{})
	//s["born"] = 0
	//s["killed"] = 0
	//s["current"] = 0
	//s["total"] = 0
	//
	//r, _ := redis.Redis.HSet(ctx, "state", s).Result()

	j, e := redis.ReJson.JSONGet("state", ".")
	if j == nil {
		t := docs.State{
			Born:    0,
			Killed:  0,
			Current: 0,
			Total:   0,
		}
		_, _ = redis.ReJson.JSONSet("state", ".", t)

		return c.JSON(t)
	}

	r := docs.State{}
	_ = json.Unmarshal(j.([]byte), &r)

	if e != nil {
		fmt.Println(e)
	}

	return c.JSON(r)
}

func GetOperations(c *fiber.Ctx) error {
	ope := make([]docs.Operation, 0)

	//opts := options.Find().SetLimit(10)
	now := time.Now()
	t := now.Add(time.Hour * 24 * -30).Format(time.RFC3339)
	f := M{"time": M{"$gt": t}}
	cursor, _ := mongo.Operations.Find(c.Context(), f)

	e := cursor.All(c.Context(), &ope)

	if e != nil {
		fmt.Println(e)
	}

	return c.JSON(ope)
}

func InsertOperation(c *fiber.Ctx) error {
	name := c.GetRespHeader("User-Name")

	o := docs.Operation{
		Executor: name,
		Time:     time.Now().Format(time.RFC3339),
	}

	_ = c.BodyParser(&o)

	j, _ := redis.ReJson.JSONGet("state", ".")

	r := docs.State{}
	_ = json.Unmarshal(j.([]byte), &r)

	switch o.Type {
	case IncrBorn:
		r.Born += *o.Quantity
		r.Current += *o.Quantity
		r.Total += *o.Quantity
	case IncrKilled:
		r.Killed += *o.Quantity
		r.Current -= *o.Quantity
		r.Total += *o.Quantity
	case MutateCurrent:
		r.Current = *o.Current
	case MutateTotal:
		r.Total = *o.Current
	}

	_, _ = redis.ReJson.JSONSet("state", ".", r)

	mr, _ := mongo.Operations.InsertOne(c.Context(), o)

	return c.JSON(mr)
}
