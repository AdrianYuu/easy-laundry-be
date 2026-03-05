package route

import (
	"github.com/AdrianYuu/easy-laundry-be/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Fiber          *fiber.App
	UserController *http.UserController
	AuthMiddleware fiber.Handler
}

func (c *Config) Setup() {
	c.Fiber.Get("/api/ping", c.UserController.Ping)

	c.Fiber.Get("/api/protected", c.AuthMiddleware, c.UserController.Ping)
}
