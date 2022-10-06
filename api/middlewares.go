package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"killifish/config"
)

var auth = func(ctx *fiber.Ctx) error {
	p := ctx.Path()
	if p != "/api/signup" && p != "/api/signin" {
		raw := ctx.GetReqHeaders()[fiber.HeaderAuthorization]

		if len(raw) < 7 {
			return ctx.Status(fiber.StatusNonAuthoritativeInformation).JSON("未携带token")
		}
		raw = raw[7:]
		token, _ := jwt.Parse(
			raw, func(token *jwt.Token) (interface{}, error) {
				return config.Secret, nil
			},
		)

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			return ctx.Status(fiber.StatusNonAuthoritativeInformation).JSON("错误的token")
		}

		ctx.Set("User-Name", claims["Name"].(string))
	}

	return ctx.Next()
}
