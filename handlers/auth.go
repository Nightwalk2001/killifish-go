package handlers

import (
	"github.com/gofiber/fiber/v2"
	gomongo "go.mongodb.org/mongo-driver/mongo"
	"killifish/docs"
	"killifish/mongo"
)

type RequestBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Signin(c *fiber.Ctx) error {
	rq := RequestBody{}
	_ = c.BodyParser(&rq)

	f := M{"_id": rq.Name, "password": rq.Password}

	p := docs.Person{}
	err := mongo.Persons.FindOne(c.Context(), f).Decode(&p)
	if err != nil {
		return c.JSON("incorrect")
	}
	token := Issue(&p)

	return c.JSON(
		fiber.Map{
			"name":      p.Name,
			"isManager": p.IsManager,
			"token":     token,
		},
	)
}

func Signup(c *fiber.Ctx) error {
	p := docs.Person{}
	_ = c.BodyParser(&p)

	_, e := mongo.Persons.InsertOne(c.Context(), p)

	if e != nil {
		if gomongo.IsDuplicateKeyError(e) {
			return c.JSON("exist")
		}
	}

	token := Issue(&p)

	return c.JSON(
		fiber.Map{
			"name":      p.Name,
			"isManager": p.IsManager,
			"token":     token,
		},
	)
}
