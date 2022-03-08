package handler

import (
	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"

	"github.com/edlorenzo/users-api/utils"
)

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")
	usersJWTMiddleware := jwt.New(
		jwt.Config{
			SigningKey: utils.JWTSecret,
			AuthScheme: "Token",
			Filter: func(c *fiber.Ctx) bool {
				if c.Method() == "GET" && c.Path() != "/api/users/feed" {
					return true
				}
				return false
			},
		})

	users := v1.Group("/users", usersJWTMiddleware)
	users.Get("/list", h.UserList)
}
