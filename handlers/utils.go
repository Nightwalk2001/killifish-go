package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type D = bson.D
type M = bson.M

var InternalServerError = fiber.StatusInternalServerError

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error error       `json:"error"`
}
