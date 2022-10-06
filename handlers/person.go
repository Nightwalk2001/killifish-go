package handlers

import (
	"github.com/gofiber/fiber/v2"
	"killifish/docs"
	"killifish/mongo"
)

func GetPersonNames(c *fiber.Ctx) error {
	ctx := c.Context()

	cursor, _ := mongo.Persons.Find(ctx, M{})

	names := make([]string, 0)

	for cursor.Next(ctx) {
		p := docs.Person{}
		_ = cursor.Decode(&p)
		names = append(names, p.Name)
	}

	return c.JSON(names)
}

func GetPersons(c *fiber.Ctx) error {
	ctx := c.Context()
	persons := make([]docs.Person, 0)

	cursor, _ := mongo.Persons.Find(ctx, M{})

	if e := cursor.All(ctx, &persons); e != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("e2")
	}

	return c.JSON(persons)
}
